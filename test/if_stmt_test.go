package test

import (
	"reflect"
	"testing"

	"github.com/wsx864321/sweet_dsl/engine"
)

const IfStmtScripts = `
package main

func main() {
	if a == 1 {
		if b == 6.0 {
			return 1
		}

		return 10
	} else {
		return 3
	}

	return 2
}
`

func TestIfStmt(t *testing.T) {
	inputPar := map[string]interface{}{
		"a": 12,
		"b": 6.0,
	}

	e := engine.NewEngine(IfStmtScripts)
	res, _ := e.Run(inputPar)

	if _, ok := res.(int64); !ok {
		t.Error(reflect.TypeOf(res), res)
		return
	}

	if res.(int64) != 3 {
		t.Error(reflect.TypeOf(res), res)
		return
	}
}
