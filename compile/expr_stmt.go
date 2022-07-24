package compile

import (
	"fmt"
	"go/ast"
)

type exprStmt struct {
	*baseStmt
}

func NewExprStmt() dslStmt {
	return &exprStmt{
		NewBaseStmt(),
	}
}

func (e *exprStmt) compileExprStmt(ctx *dataCtx, stmt *ast.ExprStmt) interface{} {
	switch stmt.X.(type) {
	case *ast.CallExpr:
		return NewExpr().compileExpr(ctx, e, stmt.X)
	default:
		panic(fmt.Sprintf("syntax error:exprStmt only support callExpr,pos:%v", ctx.fSet.Position(stmt.Pos())))
	}

}
