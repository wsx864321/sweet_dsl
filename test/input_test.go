package test

import (
	"testing"

	"github.com/wsx864321/sweet_dsl/engine"
)

const InputScripts1 = `
package main

func main() {
	for _,label := range inputData {
			if label["label"] == 1 {
				Println(label)
				return 1
			}

			if label["label"] == 2 {
				Println(label)
				return 2
			}
	}

	return 3
}
`

func TestInputTest(t *testing.T) {
	tables := []struct {
		input map[string]interface{}
		res   int64
	}{
		{
			input: map[string]interface{}{
				"inputData": []map[string]interface{}{
					{
						"class":   "game_begin",
						"label":   1,
						"content": "xxxxx",
						"score":   70,
					},
				},
			},
			res: 1,
		},
		{
			input: map[string]interface{}{
				"inputData": []map[string]interface{}{
					{
						"class":   "game_over",
						"label":   2,
						"content": "asdasdas",
						"score":   70,
					},
				},
			},
			res: 2,
		},
		{
			input: map[string]interface{}{
				"inputData": []map[string]interface{}{
					{
						"class":   "game_type",
						"label":   3,
						"content": "qqqqqq",
						"score":   80,
					},
				},
			},
			res: 3,
		},
	}
	e := engine.NewEngine(InputScripts1)
	for _, item := range tables {
		res, err := e.Run(item.input)
		if err != nil || res.(int64) != item.res {
			t.Errorf("item:%v,res:%v", item, res)
		}
	}

}
