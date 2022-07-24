package compile

import (
	"fmt"
	"go/ast"
)

type dslStmt interface {
	setParentStmt(ctx *dataCtx, stmt dslStmt)
	getParentStmt(ctx *dataCtx) dslStmt
	getVal(ctx *dataCtx, key string) (interface{}, error)
	setVal(ctx *dataCtx, key string, val interface{})
	getRuntimeVal(ctx *dataCtx, key string) (interface{}, bool)
	setRuntimeVal(ctx *dataCtx, key string, val interface{})
}

type baseStmt struct {
	parentStmt     dslStmt
	runtimeNodeCtx map[string]interface{} // 运行时作用域的上下文，非全局
}

func NewBaseStmt() *baseStmt {
	return &baseStmt{
		parentStmt:     nil,
		runtimeNodeCtx: make(map[string]interface{}),
	}
}

func (b *baseStmt) setParentStmt(ctx *dataCtx, stmt dslStmt) {
	b.parentStmt = stmt
}

func (b *baseStmt) getParentStmt(ctx *dataCtx) dslStmt {
	return b.parentStmt
}

func (b *baseStmt) getVal(ctx *dataCtx, key string) (interface{}, error) {
	val, ok := b.getRuntimeVal(ctx, key)
	if ok {
		return val, nil
	}
	// 先从runtime中的上下文获取
	parentStmt := b.parentStmt
	for parentStmt != nil {
		val, ok = parentStmt.getRuntimeVal(ctx, key)
		if ok {
			return val, nil
		}
		parentStmt = parentStmt.getParentStmt(ctx)
	}

	return ctx.getVal(key)
}

func (b *baseStmt) getRuntimeVal(ctx *dataCtx, key string) (interface{}, bool) {
	if val, ok := b.runtimeNodeCtx[key]; ok {
		return val, true
	}

	return nil, false
}

func (b *baseStmt) setVal(ctx *dataCtx, key string, val interface{}) {
	parentStmt := b.parentStmt
	if parentStmt == nil {
		b.setRuntimeVal(ctx, key, val)
	}

	for parentStmt != nil {
		_, ok := parentStmt.getRuntimeVal(ctx, key)
		if ok {
			parentStmt.setRuntimeVal(ctx, key, val)
			return
		}
		parentStmt = parentStmt.getParentStmt(ctx)
	}

	b.parentStmt.setRuntimeVal(ctx, key, val)
}

func (b *baseStmt) setRuntimeVal(ctx *dataCtx, key string, val interface{}) {
	b.runtimeNodeCtx[key] = val
}

func (b *baseStmt) compileBaseStmt(ctx *dataCtx, stmt ast.Stmt) (interface{}, bool) {
	var s dslStmt
	switch stmt := stmt.(type) {
	case *ast.IfStmt:
		s = NewIfStmt()
		s.setParentStmt(ctx, b)
		return s.(*ifStmt).compileIfStmt(ctx, stmt)
	case *ast.ReturnStmt:
		s = NewReturnStmt()
		s.setParentStmt(ctx, b)
		return s.(*returnStmt).compileReturnStmt(ctx, stmt), true
	case *ast.BlockStmt:
		s = NewBlockStmt()
		s.setParentStmt(ctx, b)
		return s.(*blockStmt).compileBlockStmt(ctx, stmt)
	case *ast.AssignStmt:
		s = NewAssignStmt()
		s.setParentStmt(ctx, b)
		s.(*assignStmt).compileAssignStmt(ctx, stmt)
	case *ast.ExprStmt:
		s = NewExprStmt()
		s.setParentStmt(ctx, b)
		return s.(*exprStmt).compileExprStmt(ctx, stmt), false
	case *ast.RangeStmt:
		s = NewRangeStmt()
		s.setParentStmt(ctx, b)
		return s.(*rangeStmt).compileRangeStmt(ctx, stmt)
	case *ast.ForStmt:
		s = NewForStmt()
		s.setParentStmt(ctx, b)
		return s.(*forStmt).compileForStmt(ctx, stmt)
	case *ast.IncDecStmt:
		s = NewIncDecStmt()
		s.setParentStmt(ctx, b)
		s.(*incDecStmt).compileIncDesStmt(ctx, stmt)
	default:
		panic(fmt.Sprintf("unsupported syntax pos:%v", ctx.fSet.Position(stmt.Pos())))
	}

	return nil, false
}
