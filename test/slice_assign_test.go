package test

import (
	"testing"

	"github.com/wsx864321/sweet_dsl/engine"
)

const SliceAssignScripts1 = `
package main

func main() {
	s := [][]int{[]int{1,2,3},[]int{4,5,6}}
	Println(s)
	Println("==========")
	s[0] = []int{41,1,231}
	Println(s)
	Println("==========")
	s[0][1] = 100
	Println(s)
}
`

func TestSliceAssign(t *testing.T) {
	e := engine.NewEngine(SliceAssignScripts1)
	e.Run(nil)
}
