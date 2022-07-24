package compile

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"strconv"
	"strings"
)

type callExpr struct {
	funcName string
	args     []ast.Expr
}

func NewCallExpr(name string, args []ast.Expr) *callExpr {
	return &callExpr{
		funcName: name,
		args:     args,
	}
}

// Invoke 暂时只支持这些函数，其他的函数可以慢慢支持，或者后续也可以直接支持系统库函数
// TODO 可以优化，利用反射去做，这样只用实现函数逻辑就行，不用每次对函数的入参进行处理
func (c *callExpr) Invoke(ctx *dataCtx, dslStmt dslStmt) interface{} {
	expr := NewExpr()
	switch c.funcName {
	case "VersionCompare":
		if len(c.args) != 3 {
			panic("syntax error:function version_compare args is must 3")
		}

		args := make([]string, 0, 3)
		for _, v := range c.args {
			arg := expr.compileExpr(ctx, dslStmt, v)
			if _, ok := arg.(string); !ok {
				panic("syntax error:arg type is error")
			}
			args = append(args, arg.(string))
		}

		return VersionCompare(args[0], args[1], args[2])
	case "Println":
		args := make([]interface{}, 0)
		for _, v := range c.args {
			args = append(args, expr.compileExpr(ctx, dslStmt, v))
		}

		Println(args...)
	case "JsonEncode":
		return JsonEncode(expr.compileExpr(ctx, dslStmt, c.args[0]))
	case "JsonDecode":
		str, _ := expr.compileExpr(ctx, dslStmt, c.args[0]).(string)
		return JsonDecode(str)
	default:
		panic(fmt.Sprintf("syntax error:unsupported function %v", c.funcName))
	}

	return nil
}

func Compare(v1, v2 string) int {
	str1List := strings.Split(v1, ".")
	str2List := strings.Split(v2, ".")
	len1 := len(str1List)
	len2 := len(str2List)
	max := len1
	if len1 < len2 {
		max = len2
	}

	for i := 0; i < max; i++ {
		v1 := uint64(0)
		if i < len1 {
			var err error
			v1, err = strconv.ParseUint(str1List[i], 10, 64)
			if err != nil {
				continue
			}
		}

		v2 := uint64(0)
		if i < len2 {
			var err error
			v2, err = strconv.ParseUint(str2List[i], 10, 64)
			if err != nil {
				continue
			}
		}

		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
	}
	return 0
}

// VersionCompare 版本号比较
func VersionCompare(v1, operator, v2 string) bool {
	com := Compare(v1, v2)
	switch operator {
	case "==", "=":
		if com == 0 {
			return true
		}
	case "<":
		if com == 2 {
			return true
		}
	case ">":
		if com == 1 {
			return true
		}
	case "<=":
		if com == 0 || com == 2 {
			return true
		}
	case ">=":
		if com == 0 || com == 1 {
			return true
		}
	}
	return false
}

func Println(a ...interface{}) {
	fmt.Println(a...)
}

func JsonEncode(par interface{}) string {
	raw, _ := json.Marshal(par)
	return string(raw)
}

func JsonDecode(s string) interface{} {
	var res interface{}
	json.Unmarshal([]byte(s), &res)

	return res
}
