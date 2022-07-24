package compile

import (
	"go/ast"

	"github.com/spf13/cast"
)

type ifStmt struct {
	*baseStmt
}

func NewIfStmt() dslStmt {
	return &ifStmt{
		NewBaseStmt(),
	}
}

func (i *ifStmt) compileIfStmt(ctx *dataCtx, a *ast.IfStmt) (interface{}, bool) {
	if a.Init != nil {

	} else {
		e := NewExpr()
		res := e.compileExpr(ctx, i, a.Cond)
		if cast.ToBool(res) {
			return i.compileBaseStmt(ctx, a.Body)
		}

		if a.Else != nil {
			return i.compileBaseStmt(ctx, a.Else)
		}
	}

	return nil, false
}
