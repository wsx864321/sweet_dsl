package compile

import (
	"fmt"
	"go/ast"

	"github.com/spf13/cast"
)

type assignStmt struct {
	*baseStmt
}

func NewAssignStmt() dslStmt {
	return &assignStmt{
		NewBaseStmt(),
	}
}

func (a *assignStmt) compileAssignStmt(ctx *dataCtx, stmt *ast.AssignStmt) {
	if len(stmt.Lhs) != len(stmt.Rhs) {
		panic(fmt.Sprintf("syntax error: bad assignStmt nums，pos:%v", ctx.fSet.Position(stmt.Pos())))
	}

	for i, lhs := range stmt.Lhs {
		val := NewExpr().compileExpr(ctx, a, stmt.Rhs[i])
		switch lh := lhs.(type) {
		case *ast.Ident:
			a.setVal(ctx, lh.Name, val)
		case *ast.IndexExpr:
			a.assignIndexExpr(ctx, stmt, lh, val)
		}
	}
}

func (a *assignStmt) assignIndexExpr(ctx *dataCtx, stmt *ast.AssignStmt, expr ast.Expr, origin interface{}) {
	switch expr := expr.(type) {
	case *ast.IndexExpr:
		src := NewExpr().compileExpr(ctx, a, expr.X)
		idx := NewExpr().compileExpr(ctx, a, expr.Index)
		switch src := src.(type) {
		case []interface{}:
			// 这个地方到底要不要这么设计呢？？？？？
			if origin != nil && src[cast.ToInt(idx)] != nil && (NewBasicLit().GetDataType(src[cast.ToInt(idx)]) != (NewBasicLit().GetDataType(origin))) {
				panic(fmt.Sprintf("inconsistent assignment type pos:%v", ctx.fSet.Position(stmt.Pos())))
			}
			src[cast.ToInt(idx)] = origin
			a.assignIndexExpr(ctx, stmt, expr.X, src)
		case map[interface{}]interface{}:
			// 这个地方到底要不要这么设计呢？？？？？
			if origin != nil && src[idx] != nil && (NewBasicLit().GetDataType(src[idx]) != (NewBasicLit().GetDataType(origin))) {
				panic(fmt.Sprintf("Inconsistent assignment type pos:%v", ctx.fSet.Position(stmt.Pos())))
			}
			src[idx] = origin
			a.assignIndexExpr(ctx, stmt, expr.X, src)
		case map[string]interface{}:
			// 这个地方到底要不要这么设计呢？？？？？
			if origin != nil && src[cast.ToString(idx)] != nil && (NewBasicLit().GetDataType(src[cast.ToString(idx)]) != (NewBasicLit().GetDataType(origin))) {
				panic(fmt.Sprintf("Inconsistent assignment type pos:%v", ctx.fSet.Position(stmt.Pos())))
			}
			src[cast.ToString(idx)] = origin
			a.assignIndexExpr(ctx, stmt, expr.X, src)
		default:
			panic(fmt.Sprintf("key is not exsit pos:%v", ctx.fSet.Position(stmt.Pos())))
		}
	case *ast.Ident:
		a.setVal(ctx, expr.Name, origin)
	}

}
