package test

import (
	"fmt"
	"testing"

	"github.com/wsx864321/sweet_dsl/engine"
)

const MapExprScripts1 = `
package main

func main() {
	s := map[string]interface{}{
		"name" : "wsx",

	}
	Println(s["name"])
}
`

const MapExprScripts2 = `
package main

func main() {
	s := map[string]interface{}{
		"a":map[string]interface{}{
			"name":"wsx",
		},

	}
	Println(s["a"])
	Println(s["a"]["name"])
}
`

const MapExprScripts3 = `
package main

func main() {
	s := map[string]interface{}{
		"name":"wsx",
		"age":25,
	}

	for k,v := range s {
		Println(k, v)
	}

	for k := range s {
		Println(k)
	}
}
`

func TestMapExpr(t *testing.T) {
	e := engine.NewEngine(MapExprScripts1)
	e.Run(nil)
	fmt.Println("=============")
	e1 := engine.NewEngine(MapExprScripts2)
	e1.Run(nil)
	fmt.Println("=============")
	e2 := engine.NewEngine(MapExprScripts3)
	e2.Run(nil)
}
