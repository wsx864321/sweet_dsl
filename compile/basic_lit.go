package compile

import (
	"fmt"
	"reflect"
)

type Datatype uint16

const (
	TypeInt       Datatype = iota // int
	TypeFloat                     // float
	TypeString                    // string
	TypeBool                      // bool
	TypeArray                     // 数组
	TypeInterface                 // 接口
	TypeMap                       // 映射
	TypePtr                       // 指针
	TypeStruct                    // 结构体，暂时不支持
	TypeSlice                     // 切片
)

type BasicLit struct {
}

func NewBasicLit() *BasicLit {
	return &BasicLit{}
}

func (b *BasicLit) GetDataType(val interface{}) Datatype {
	t := reflect.TypeOf(val).Kind()
	switch t {
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return TypeInt
	case reflect.Float32, reflect.Float64:
		return TypeFloat
	case reflect.String:
		return TypeString
	case reflect.Bool:
		return TypeBool
	case reflect.Array:
		return TypeArray
	case reflect.Interface:
		return TypeInterface
	case reflect.Map:
		return TypeMap
	case reflect.Slice:
		return TypeSlice
	case reflect.Ptr:
		return TypePtr
	default:
		panic(fmt.Sprintf("syntax error:unsupported data type,val:%v", val))
	}

}
