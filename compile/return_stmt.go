package compile

import (
	"fmt"
	"go/ast"
)

type returnStmt struct {
	*baseStmt
}

func NewReturnStmt() *returnStmt {
	return &returnStmt{
		NewBaseStmt(),
	}
}

func (r *returnStmt) compileReturnStmt(ctx *dataCtx, stmt *ast.ReturnStmt) interface{} {
	if len(stmt.Results) != 1 {
		panic(fmt.Sprintf("syntax error:multiple return values are not supported,pos:%v", ctx.fSet.Position(stmt.Pos())))
	}

	e := NewExpr()

	return e.compileExpr(ctx, r, stmt.Results[0])
}
