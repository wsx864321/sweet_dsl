package test

import (
	"testing"

	"github.com/wsx864321/sweet_dsl/engine"
)

const MapAssignScripts1 = `
package main

func main() {
	s := map[string]interface{}{
		"name":"wsx",
		"age":25,
		"detail_addr":map[string]interface{}{
			"city":"xxxxx",
		},
	}
	s["name"] = "zmq"
	s["detail_addr"]["city"] = "wuhan"
	Println(s)
}
`

func TestMapAssign(t *testing.T) {
	e := engine.NewEngine(MapAssignScripts1)
	e.Run(nil)
}
