package compile

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"strconv"

	"github.com/spf13/cast"
)

type expr struct{}

func NewExpr() *expr {
	return &expr{}
}

func (e *expr) compileExpr(ctx *dataCtx, dslStmt dslStmt, expr ast.Expr) interface{} {
	switch x := expr.(type) {
	case *ast.BasicLit:
		return e.lit2Value(ctx, x)
	case *ast.BinaryExpr:
		left := e.compileExpr(ctx, dslStmt, x.X)
		right := e.compileExpr(ctx, dslStmt, x.Y)
		return e.compileBinaryExpr(ctx, left, right, x.Op)
	case *ast.CallExpr:
		funcName := x.Fun.(*ast.Ident).Name
		return NewCallExpr(funcName, x.Args).Invoke(ctx, dslStmt)
	case *ast.ParenExpr:
		return e.compileExpr(ctx, dslStmt, x.X)
	case *ast.Ident:
		return e.compileIdentExpr(ctx, dslStmt, x)
	case *ast.CompositeLit:
		switch x.Type.(type) {
		case *ast.ArrayType:
			return e.compileArrayType(ctx, dslStmt, x)
		case *ast.MapType:
			return e.compileMapType(ctx, dslStmt, x)
		}
	case *ast.IndexExpr:
		return e.compileIndexExpr(ctx, dslStmt, x)
	// case *ast.SliceExpr: // TODO,expr[low:mid:high] 表达式，暂时先不支持,看之后有没有类似需求，有的话再支持下

	default:
		panic(fmt.Sprintf("syntax error:unsupported expr,pos:%v", ctx.fSet.Position(expr.Pos())))

	}

	return nil
}

func (e *expr) lit2Value(ctx *dataCtx, lit *ast.BasicLit) interface{} {
	var res interface{}
	switch lit.Kind {
	case token.INT:
		return cast.ToInt64(lit.Value)
	case token.STRING:
		res, _ := strconv.Unquote(lit.Value)
		return res
	case token.FLOAT:
		return cast.ToFloat64(lit.Value)
	case token.CHAR:
		return lit.Value
	}

	return res
}

func (e *expr) compileIdentExpr(ctx *dataCtx, dslStmt dslStmt, expr *ast.Ident) interface{} {
	if expr.Name == "true" {
		return true
	} else if expr.Name == "false" {
		return false
	} else if expr.Name == "string" {
		return "string"
	}

	res, err := dslStmt.getVal(ctx, expr.Name)
	if err != nil {
		panic(err)
	}

	return res
}

func (e *expr) compileArrayType(ctx *dataCtx, dslStmt dslStmt, expr *ast.CompositeLit) interface{} {
	res := make([]interface{}, 0)
	for _, elt := range expr.Elts {
		res = append(res, e.compileExpr(ctx, dslStmt, elt))
	}

	return res
}

func (e *expr) compileMapType(ctx *dataCtx, dslStmt dslStmt, expr *ast.CompositeLit) interface{} {
	res := make(map[interface{}]interface{})
	for _, item := range expr.Elts {
		key := NewExpr().compileExpr(ctx, dslStmt, item.(*ast.KeyValueExpr).Key)
		val := NewExpr().compileExpr(ctx, dslStmt, item.(*ast.KeyValueExpr).Value)
		res[key] = val
	}

	return res
}

func (e *expr) compileIndexExpr(ctx *dataCtx, dslStmt dslStmt, expr *ast.IndexExpr) interface{} {
	src := e.compileExpr(ctx, dslStmt, expr.X)
	idx := e.compileExpr(ctx, dslStmt, expr.Index)
	switch src := src.(type) {
	case map[interface{}]interface{}:
		return src[idx]
	case []interface{}:
		return src[cast.ToInt(idx)]
	case map[string]interface{}:
		return src[cast.ToString(idx)]
	}
	return nil
}

func (e *expr) compileBinaryExpr(ctx *dataCtx, left, right interface{}, op token.Token) interface{} {
	var res interface{}
	switch op {
	case token.EQL: // ==
		return e.eqlExpr(ctx, left, right)
	case token.NEQ:
		return e.neqExpr(ctx, left, right)
	case token.LSS: // <
		return e.lssExpr(ctx, left, right)
	case token.GTR: // >
		return e.gtrExpr(ctx, left, right)
	case token.GEQ: // >=
		return e.geqExpr(ctx, left, right)
	case token.LEQ: // <=
		return e.leqExpr(ctx, left, right)
	case token.LAND: // &&
		return e.landExpr(ctx, left, right)
	case token.LOR: // ||
		return e.lorExpr(ctx, left, right)
	case token.ADD: // +
		return e.addExpr(ctx, left, right)
	case token.SUB: // -
		return e.subExpr(ctx, left, right)
	case token.MUL: // *
		return e.mulExpr(ctx, left, right)
	case token.QUO: // -
		return e.quoExpr(ctx, left, right)
	case token.REM: // %
		return e.remExpr(ctx, left, right)
	}

	return res
}

func (e *expr) neqExpr(ctx *dataCtx, left, right interface{}) interface{} {
	var res interface{}
	b := NewBasicLit()
	t := b.GetDataType(left)
	if t != b.GetDataType(right) {
		panic("syntax error: mismatched type")
	}

	switch t {
	case TypeInt:
		return cast.ToInt64(left) != cast.ToInt64(right)
	case TypeFloat:
		return cast.ToFloat64(left) != cast.ToFloat64(right)
	case TypeString:
		return cast.ToString(left) != cast.ToString(right)
	default:
		panic(fmt.Sprintf("syntax error: unsupported data type"))
	}

	return res
}

func (e *expr) eqlExpr(ctx *dataCtx, left, right interface{}) interface{} {
	var res interface{}
	b := NewBasicLit()
	t := b.GetDataType(left)
	if t != b.GetDataType(right) {
		panic("syntax error: mismatched type")
	}

	switch t {
	case TypeInt:
		return cast.ToInt64(left) == cast.ToInt64(right)
	case TypeFloat:
		return cast.ToFloat64(left) == cast.ToFloat64(right)
	case TypeString:
		return cast.ToString(left) == cast.ToString(right)
	default:
		panic(fmt.Sprintf("syntax error: unsupported data type"))
	}

	return res
}

func (e *expr) lssExpr(ctx *dataCtx, left, right interface{}) interface{} {
	var res interface{}
	b := NewBasicLit()
	t := b.GetDataType(left)
	if t != b.GetDataType(right) {
		panic("syntax error: mismatched type")
	}

	switch t {
	case TypeInt:
		return cast.ToInt64(left) < cast.ToInt64(right)
	case TypeFloat:
		return cast.ToFloat64(left) < cast.ToFloat64(right)
	case TypeString:
		return cast.ToString(left) < cast.ToString(right)
	default:
		panic("syntax error: unsupported data type")
	}

	return res
}

func (e *expr) gtrExpr(ctx *dataCtx, left, right interface{}) interface{} {
	var res interface{}
	b := NewBasicLit()
	t := b.GetDataType(left)
	if t != b.GetDataType(right) {
		panic("syntax error: mismatched type")
	}

	switch t {
	case TypeInt:
		return cast.ToInt64(left) > cast.ToInt64(right)
	case TypeFloat:
		return cast.ToFloat64(left) > cast.ToFloat64(right)
	case TypeString:
		return cast.ToString(left) > cast.ToString(right)
	default:
		panic("syntax error: unsupported data type")
	}

	return res
}

func (e *expr) geqExpr(ctx *dataCtx, left, right interface{}) interface{} {
	var res interface{}
	b := NewBasicLit()
	t := b.GetDataType(left)
	if t != b.GetDataType(right) {
		panic("syntax error: mismatched type")
	}

	switch t {
	case TypeInt:
		return cast.ToInt64(left) >= cast.ToInt64(right)
	case TypeFloat:
		return cast.ToFloat64(left) >= cast.ToFloat64(right)
	case TypeString:
		return cast.ToString(left) >= cast.ToString(right)
	default:
		panic("syntax error: unsupported data type")
	}

	return res
}

func (e *expr) leqExpr(ctx *dataCtx, left, right interface{}) interface{} {
	var res interface{}
	b := NewBasicLit()
	t := b.GetDataType(left)
	if t != b.GetDataType(right) {
		panic("syntax error: mismatched type")
	}

	switch t {
	case TypeInt:
		return cast.ToInt64(left) <= cast.ToInt64(right)
	case TypeFloat:
		return cast.ToFloat64(left) <= cast.ToFloat64(right)
	case TypeString:
		return cast.ToString(left) <= cast.ToString(right)
	default:
		panic("syntax error: unsupported data type")
	}

	return res
}

func (e *expr) landExpr(ctx *dataCtx, left, right interface{}) interface{} {
	t := NewBasicLit()
	if t.GetDataType(left) != TypeBool {
		panic(fmt.Sprintf("syntax error: bad binary type= %v \n", reflect.TypeOf(left)))
	}

	if t.GetDataType(right) != TypeBool {
		panic(fmt.Sprintf("syntax error: bad binary type= %v \n", reflect.TypeOf(right)))
	}

	return cast.ToBool(left) && cast.ToBool(right)
}

func (e *expr) lorExpr(ctx *dataCtx, left, right interface{}) interface{} {
	t := NewBasicLit()
	if t.GetDataType(left) != TypeBool {
		panic(fmt.Sprintf("syntax error: bad binary type= %v \n", reflect.TypeOf(left)))
	}

	if t.GetDataType(right) != TypeBool {
		panic(fmt.Sprintf("syntax error: bad binary type= %v \n", reflect.TypeOf(right)))
	}

	return cast.ToBool(left) || cast.ToBool(right)
}

func (e *expr) addExpr(ctx *dataCtx, left, right interface{}) interface{} {
	var res interface{}
	switch reflect.TypeOf(left).Kind() {
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return cast.ToInt64(left) + cast.ToInt64(right)
	case reflect.Float32, reflect.Float64:
		return cast.ToFloat64(left) + cast.ToFloat64(right)
	default:
		panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", reflect.TypeOf(left)))
	}
	return res
}

func (e *expr) subExpr(ctx *dataCtx, left, right interface{}) interface{} {
	var res interface{}
	switch reflect.TypeOf(left).Kind() {
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return cast.ToInt64(left) - cast.ToInt64(right)
	case reflect.Float32, reflect.Float64:
		return cast.ToFloat64(left) - cast.ToFloat64(right)
	default:
		panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", reflect.TypeOf(left)))

	}
	return res
}

func (e *expr) mulExpr(ctx *dataCtx, left, right interface{}) interface{} {
	var res interface{}
	switch reflect.TypeOf(left).Kind() {
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return cast.ToInt64(left) * cast.ToInt64(right)
	case reflect.Float32, reflect.Float64:
		return cast.ToFloat64(left) * cast.ToFloat64(right)
	default:
		panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", reflect.TypeOf(left)))

	}
	return res
}

func (e *expr) quoExpr(ctx *dataCtx, left, right interface{}) interface{} {
	var res interface{}
	switch reflect.TypeOf(left).Kind() {
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return cast.ToInt64(left) / cast.ToInt64(right)
	case reflect.Float32, reflect.Float64:
		return cast.ToFloat64(left) / cast.ToFloat64(right)
	default:
		panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", reflect.TypeOf(left)))

	}
	return res
}

func (e *expr) remExpr(ctx *dataCtx, left, right interface{}) interface{} {
	var res interface{}
	switch reflect.TypeOf(left).Kind() {
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return cast.ToInt64(left) % cast.ToInt64(right)
	default:
		panic(fmt.Sprintf("syntax error: bad binary type= %#v \n", reflect.TypeOf(left)))

	}
	return res
}
