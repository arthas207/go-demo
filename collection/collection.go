package collection

import "fmt"

type person struct {
	name string
	sex  string
	age  int
	// test []int 会报错，只支持简单值比较
}

func Example() {
	var slice1 = make([]int64, 4)
	slice1[0] = 5
	slice1[1] = 6
	slice1[2] = 7
	slice1[3] = 8
	fmt.Println("slice1:", slice1[:2])
	var userMap = map[string]int{
		"Bob":   5,
		"Alice": 4,
	}
	fmt.Println("userMap", userMap)
	fmt.Println("userMap get Bob", userMap["Bob"])
	var slice2 = []int{4, 5, 7, 8}
	fmt.Println("slice2:", slice2[2:])
	var ageMap = make(map[string]int)
	ageMap["Alice"] = 4
	fmt.Println("ageMap:", ageMap)
	var person1 = person{"Tom", "male", 45}
	var personMap = make(map[person]int)
	personMap[person1] = 10
	fmt.Println("personMap:", personMap)
	// 只支持简单值比较==,深层对象嵌套会报错
	age := personMap[person{"Tom", "male", 45}]
	fmt.Println(age)
	age2 := personMap[person1]
	fmt.Sprintln(age2)
}
