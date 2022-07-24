## sweet_dsl
### 介绍
sweet_dsl是基于go/parse暴露出来的API做的一个简单的规则引擎，简单点说就是一个dsl，其语法规则完全兼容Go，只是对Go的一些不必要的语法做了裁剪。

### 语法规则
兼容大部分Go语言语法，可以认为是语法不严格的Go语言

#### 数据类型
> int、float、string、map、slice

#### 操作符
> < >  >= <= && || !=

#### 语句
> for、for range、if

#### 支持函数
> Println、VersionCompare、JsonEncode、JsonDecode
> 
> 理论上任意函数都可以进行支持

### Demo
```go
package test

import (
	"reflect"
	"sweet_dsl/engine"
	"testing"
)

const IfStmtScripts = `
package main

func main() {
	if a == 1 {
		if b == 6.0 {
			return 1
		}

		return 10
	} else {
		return 3
	}

	return 2
}
`

func TestIfStmt(t *testing.T) {
	inputPar := map[string]interface{}{
		"a": 12,
		"b": 6.0,
	}

	e := engine.NewEngin(IfStmtScripts)
	res, _ := e.Run(inputPar)

	if _, ok := res.(int64); !ok {
		t.Error(reflect.TypeOf(res), res)
		return
	}

	if res.(int64) != 3 {
		t.Error(reflect.TypeOf(res), res)
		return
	}
}
```
### 说明
> 如在使用过程中有需要支持的语法可以提issues或者直接联系我
