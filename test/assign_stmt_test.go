package test

import (
	"reflect"
	"testing"

	"github.com/wsx864321/sweet_dsl/engine"
)

const AssignStmtScripts = `
package main

func main() {
	b := 6.0
	if a == 1 {
		if b == 6.0 {
			return 1
		}

		return 10
	} else if a == 12{
		return 3
	}

	return 2
}
`

func TestAssignStmt(t *testing.T) {
	inputPar := map[string]interface{}{
		"a": 1,
	}

	e := engine.NewEngine(AssignStmtScripts)
	res, _ := e.Run(inputPar)
	t.Log(res)
	if _, ok := res.(int64); !ok {
		t.Error(reflect.TypeOf(res), res)
		return
	}

	if res.(int64) != 1 {
		t.Error(reflect.TypeOf(res), res)
		return
	}
}
