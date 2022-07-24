package compile

import (
	"fmt"
	"go/ast"
)

type forStmt struct {
	*baseStmt
}

func NewForStmt() dslStmt {
	return &forStmt{
		NewBaseStmt(),
	}
}

func (f *forStmt) compileForStmt(ctx *dataCtx, stmt *ast.ForStmt) (interface{}, bool) {
	if stmt.Init != nil {
		f.compileBaseStmt(ctx, stmt.Init)
	}

	if stmt.Cond != nil {
		for {
			b := NewExpr().compileExpr(ctx, f, stmt.Cond)
			if b, ok := b.(bool); ok {
				if b {
					res, isEnd := f.compileBaseStmt(ctx, stmt.Body)
					if isEnd {
						return res, true
					}

					f.compileBaseStmt(ctx, stmt.Post)
				} else {
					break
				}
			} else {
				panic(fmt.Sprintf("unsupported syntax pos:%v", ctx.fSet.Position(stmt.Pos())))
			}
		}

	}

	return nil, false
}
