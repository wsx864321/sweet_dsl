package test

import (
	"testing"

	"github.com/wsx864321/sweet_dsl/engine"
)

const JsonScripts1 = `
package main

func main() {
	jsonStr := "{\"name\":\"wsx\",\"age\":17,\"character\":[{\"a\":\"aa\",\"bb\":\"bb\",\"wsx\":[{\"omg\":\"rng\"}]},{\"c\":\"ccccc\"}],\"addr\":{\"province\":\"hubei\",\"city\":\"wuhan\",\"detail\":{\"detail_addr\":\"xxxxx\"}}}\n"
	m := JsonDecode(jsonStr)
	Println(m)
	Println(m["addr"])
	Println(JsonEncode(m))
}
`

func TestJson(t *testing.T) {
	e := engine.NewEngine(JsonScripts1)
	e.Run(nil)
}
