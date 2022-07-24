package compile

import "go/ast"

type blockStmt struct {
	*baseStmt
}

func NewBlockStmt() dslStmt {
	return &blockStmt{
		NewBaseStmt(),
	}
}

func (b *blockStmt) compileBlockStmt(ctx *dataCtx, stmt *ast.BlockStmt) (interface{}, bool) {
	for _, item := range stmt.List {
		res, isEnd := b.compileBaseStmt(ctx, item)
		if isEnd {
			return res, true
		}
	}

	return nil, false
}
