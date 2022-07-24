package compile

import (
	"go/ast"
	"go/token"
)

type Compile struct {
	file *ast.File
	ctx  *dataCtx
}

func NewCompile(f *ast.File, fSet *token.FileSet, ctx map[string]interface{}) *Compile {
	dataCtx := NewDataCtx(fSet)
	dataCtx.originCtx = ctx

	return &Compile{
		file: f,
		ctx:  dataCtx,
	}
}

func (c *Compile) Compile() interface{} {
	for _, d := range c.file.Decls {
		switch d := d.(type) {
		case *ast.GenDecl:
		case *ast.FuncDecl:
			if d.Name.String() == "main" {
				return c.compileFuncDecl(d)
			} else {
				panic("syntax error: The entry point must be main function")
			}

		default:
			panic("unsupported syntax")
		}
	}

	return nil
}

func (c *Compile) compileFuncDecl(d *ast.FuncDecl) interface{} {
	base := NewBaseStmt()
	for _, stmt := range d.Body.List {
		res, isEnd := base.compileBaseStmt(c.ctx, stmt)
		if isEnd {
			return res
		}
	}

	return nil
}
