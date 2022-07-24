package compile

import (
	"go/ast"
	"reflect"
)

type rangeStmt struct {
	*baseStmt
}

func NewRangeStmt() dslStmt {
	return &rangeStmt{
		NewBaseStmt(),
	}
}

func (r *rangeStmt) compileRangeStmt(ctx *dataCtx, stmt *ast.RangeStmt) (interface{}, bool) {
	rangVal := NewExpr().compileExpr(ctx, r, stmt.Key.(*ast.Ident).Obj.Decl.(*ast.AssignStmt).Rhs[0].(*ast.UnaryExpr).X)
	kName := stmt.Key.(*ast.Ident).Name
	isExistVal := false
	vName := ""
	if len(stmt.Key.(*ast.Ident).Obj.Decl.(*ast.AssignStmt).Lhs) == 2 {
		isExistVal = true
		vName = stmt.Key.(*ast.Ident).Obj.Decl.(*ast.AssignStmt).Lhs[1].(*ast.Ident).Name
	}

	if reflect.ValueOf(rangVal).Kind() == reflect.Slice {
		var vals []interface{}
		v := reflect.ValueOf(rangVal)
		for i := 0; i < v.Len(); i++ {
			vals = append(vals, v.Index(i).Interface())
		}

		for i, v := range vals {
			if isExistVal {
				r.setVal(ctx, vName, v)
			}
			r.setVal(ctx, kName, i)
			res, isEnd := r.compileBaseStmt(ctx, stmt.Body)
			if isEnd {
				return res, true
			}
		}

		return nil, false
	}

	switch rangVal := rangVal.(type) {
	case map[interface{}]interface{}:
		for k, v := range rangVal {
			if isExistVal {
				r.setVal(ctx, vName, v)
			}
			r.setVal(ctx, kName, k)
			res, isEnd := r.compileBaseStmt(ctx, stmt.Body)
			if isEnd {
				return res, true
			}
		}
	case map[string]interface{}:
		for k, v := range rangVal {
			if isExistVal {
				r.setVal(ctx, vName, v)
			}
			r.setVal(ctx, kName, k)
			res, isEnd := r.compileBaseStmt(ctx, stmt.Body)
			if isEnd {
				return res, true
			}
		}
	}

	return nil, false
}
