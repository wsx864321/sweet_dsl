package test

import (
	"testing"

	"github.com/wsx864321/sweet_dsl/engine"
)

const SliceScripts1 = `
package main

func main() {
	s := []map[string]string{
		map[string]string{"name": "wsx"},
		map[string]string{"name": "zmq"},
	}
	Println(s)
	for _,v := range s {
		Println(v["name"])
	}
}
`

const SliceScripts2 = `
package main

func main() {
	s := [][]int{[]int{1,2,3},[]int{4,5,6}}
	Println(s[1])
	if s[0][1] == 2 {
		Println("wsx")
	}
}
`

/**
* 声明复合类型的slice时，类似[]map[string]string{}时
* 必须这样声明：
*	s := []map[string]string{
*		map[string]string{"name": "wsx"},
*		map[string]string{"name": "zmq"},
*	}
* 元素里面必须函数map[string]string
 */
func TestSlice(t *testing.T) {
	e := engine.NewEngine(SliceScripts1)
	e.Run(nil)

	e1 := engine.NewEngine(SliceScripts2)
	e1.Run(nil)
}
