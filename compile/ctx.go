package compile

import (
	"errors"
	"fmt"
	"go/token"
)

type dataCtx struct {
	fSet      *token.FileSet
	originCtx map[string]interface{} // 初始上下文
}

func NewDataCtx(f *token.FileSet) *dataCtx {
	return &dataCtx{
		fSet:      f,
		originCtx: make(map[string]interface{}),
	}
}

func (c *dataCtx) getOriginVal(key string) (interface{}, error) {

	if res, ok := c.originCtx[key]; ok {
		return res, nil
	}

	return nil, errors.New(fmt.Sprintf("variable %s not found", key))
}

func (c *dataCtx) getVal(key string) (interface{}, error) {
	if res, ok := c.originCtx[key]; ok {
		return res, nil
	}

	return nil, errors.New(fmt.Sprintf("variable %s not found", key))
}
