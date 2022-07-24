package test

import (
	"fmt"
	"testing"

	"github.com/wsx864321/sweet_dsl/engine"
)

const VersionCompareScripts = `
package main

func main() {
	if VersionCompare("6.5.0", ">", "6.4.2") && true {
		return 1
	}

	return 2
}
`

const PrintlnScripts = `
package main

func main() {
	name := "wsx"
	
	age := 25
	
	hobby := []string{"篮球","足球"}

	Println(name, age, hobby)
}
`

func TestFunction(t *testing.T) {
	e := engine.NewEngine(VersionCompareScripts)
	fmt.Println(e.Run(nil))

	e1 := engine.NewEngine(PrintlnScripts)
	e1.Run(nil)
}
