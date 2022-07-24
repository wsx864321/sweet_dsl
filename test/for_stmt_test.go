package test

import (
	"fmt"
	"testing"

	"github.com/wsx864321/sweet_dsl/engine"
)

const ForScripts1 = `
package main

func main() {
	i := 0
	for ; i < 5; i++ {
		Println(i)
	}
}
`

const ForScripts2 = `
package main

func main() {
	for i := 0; i < 5; i++ {
		Println(i)
	}
}
`

func TestForStmt(t *testing.T) {
	e := engine.NewEngine(ForScripts1)
	e.Run(nil)
	fmt.Println("===========")
	e1 := engine.NewEngine(ForScripts2)
	e1.Run(nil)
}
