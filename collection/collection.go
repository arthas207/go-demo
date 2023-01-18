package collection

import "fmt"

// 小写是内部变量和方法,大写是可以跨包访问的
type person struct {
	name string
	sex  string
	age  int
	// test []int 会报错，只支持简单值比较
}

type heap struct {
}

func Example() {
	// slice
	var slice = []int{5, 6, 7, 8}
	fmt.Println("slice first two element,third to last element", slice[:2], slice[2:])
	// map
	var userMap = map[string]int{
		"Bob":   5,
		"Alice": 4,
	}
	fmt.Println("userMap", userMap)
	fmt.Println("userMap get Bob", userMap["Bob"])
	var tom = person{"Tom", "male", 45}
	var personMap = map[person]int{
		tom: 10,
	}
	fmt.Println("personMap:", personMap)
	// map的key只支持简单值比较==,深层对象嵌套会报错
	age := personMap[person{"Tom", "male", 45}]
	age2 := personMap[tom]
	// can't get age
	fmt.Sprintf("age:%d,age2:%d", age, age2)
}
