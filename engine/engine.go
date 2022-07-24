package engine

import (
	"go/parser"
	"go/token"

	"github.com/wsx864321/sweet_dsl/compile"
)

type Engine struct {
	Source string
}

func NewEngine(src string) *Engine {
	return &Engine{
		Source: src,
	}
}

func (e *Engine) Run(ctx map[string]interface{}) (res interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", e.Source, 0)
	if err != nil {
		panic(err)
	}

	c := compile.NewCompile(f, fset, ctx)
	return c.Compile(), nil
}
