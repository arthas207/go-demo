package reflection

import (
	"fmt"
	"reflect"
)

type Pair interface {
	DoSome(val string) string
}

type keyValuePair struct {
	name  string
	value string
}

func (keyValuePair keyValuePair) DoSome(val string) string {
	fmt.Printf("do some,%s:%s\n", keyValuePair.name, val)
	return "success"
}

func Example() {
	// 相较于java较为复杂，类型和值分属于不同的方法
	var pair = keyValuePair{name: "key", value: "value"}
	fmt.Printf("type : %s\n", reflect.TypeOf(pair))
	fmt.Printf("value : %s\n", reflect.ValueOf(pair).FieldByName("name"))
	var pp = &pair
	fmt.Println("elem:", reflect.ValueOf(pp).Elem())
	var r = reflect.ValueOf(pair).MethodByName("DoSome").Call([]reflect.Value{reflect.ValueOf("test")})
	fmt.Println("method : ", r)
}
