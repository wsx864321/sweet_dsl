package compile

import (
	"fmt"
	"go/ast"
	"go/token"
)

type incDecStmt struct {
	*baseStmt
}

func NewIncDecStmt() dslStmt {
	return &incDecStmt{
		NewBaseStmt(),
	}
}

func (i *incDecStmt) compileIncDesStmt(ctx *dataCtx, stmt *ast.IncDecStmt) {
	if token.INC != stmt.Tok && token.DEC != stmt.Tok {
		panic(fmt.Sprintf("unsupported syntax pos:%v", ctx.fSet.Position(stmt.Pos())))
	}

	res := NewExpr().compileExpr(ctx, i, stmt.X)
	if token.INC == stmt.Tok {
		i.setVal(ctx, stmt.X.(*ast.Ident).Name, NewExpr().addExpr(ctx, res, 1))
	} else if token.DEC == stmt.Tok {
		i.setVal(ctx, stmt.X.(*ast.Ident).Name, NewExpr().subExpr(ctx, res, 1))
	}
}
