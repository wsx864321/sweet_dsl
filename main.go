package main

import (
	"fmt"

	"github.com/wsx864321/sweet_dsl/engine"
)

var src = `
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
	
	s1 := [][]int{[]int{1,2,3},[]int{4,5,6}}
	s1[0] = []int{1,2,3,4,5}
	Println(s1)
	
	jsonStr := "{\"name\":\"wsx\",\"age\":17,\"character\":[{\"a\":\"aa\",\"bb\":\"bb\",\"wsx\":[{\"omg\":\"rng\"}]},{\"c\":\"ccccc\"}],\"addr\":{\"province\":\"hubei\",\"city\":\"wuhan\",\"detail\":{\"detail_addr\":\"xxxxx\"}}}\n"
	m := JsonDecode(jsonStr)
	for _ ,item := range m {
		Println(item)
	}
	Println(m)
	Println(m["addr"])
	Println(JsonEncode(m))
	
	return "wsx"
}
`

// eg
func main() {
	e := engine.NewEngine(src)
	fmt.Println(e.Run(nil))
}
