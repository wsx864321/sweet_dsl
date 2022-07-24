package test

import (
	"fmt"
	"testing"

	"github.com/wsx864321/sweet_dsl/engine"
)

const RangeStmtScripts1 = `
package main

func main() {
	s := []int{1,2,3,4,5}
	for i,v := range s {
		if v == 6 {
			return i
		}
	}

	return 1
}
`

func TestRangeStmt(t *testing.T) {
	e1 := engine.NewEngine(RangeStmtScripts1)
	fmt.Println(e1.Run(nil))
}
