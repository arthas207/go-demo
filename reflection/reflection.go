package reflection

import (
	"fmt"
	"reflect"
)

type Pair interface {
	doSome() string
}

type keyValuePair struct {
	name  string
	value string
}

func (keyValuePair keyValuePair) DoSome(val string) string {
	fmt.Printf("do some,%s,%s\n", keyValuePair.name, val)
	return "success"
}

func Example() {
	var pair keyValuePair = keyValuePair{name: "key", value: "value"}
	fmt.Printf("type : %s\n", reflect.TypeOf(pair))
	fmt.Printf("value : %s\n", reflect.ValueOf(pair).FieldByName("name"))
	var pp = &pair
	fmt.Println("elem:", reflect.ValueOf(pp).Elem())
	var r = reflect.ValueOf(pair).MethodByName("DoSome").Call([]reflect.Value{reflect.ValueOf("test")})
	fmt.Println("method : ", r)
}
