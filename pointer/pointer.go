package pointer

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Person struct {
	age  int
	name string
}
type s1 struct {
	id   int
	name string
}

type s2 struct {
	field1 *[5]byte
	filed2 int
}

type deepObj struct {
	id   int
	name string
	stru s1
}

func Example() {
	n := 10
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = i
	}
	fmt.Println(b)
	// [0 1 2 3 4 5 6 7 8 9]
	// &表示变量的指针
	fmt.Println(&b[0])
	// 数组的指针不是第一个元素的地址，也不是最后一个元素
	fmt.Println(uintptr(unsafe.Pointer(&b)))
	fmt.Println(uintptr(unsafe.Pointer(&b[0])))
	fmt.Println(uintptr(unsafe.Pointer(&b[0])) + 9*unsafe.Sizeof(b[0]))

	// *表示取指针的内容,并且复制对象
	fmt.Println(*(&b[0]))
	var elem = &b[0]
	// 类型上的*代表是指针类型,比如*int代表指向int类型变量的指针
	fmt.Println(reflect.TypeOf(elem).Elem())
	// uintptr+unsafe.Pointer用来直接访问和操作内存
	fmt.Println(reflect.TypeOf(uintptr(unsafe.Pointer(elem))))
	// 取slice的最后的一个元素 uintptr+offset可以转换成unsafe.Pointer
	end := unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + 9*unsafe.Sizeof(b[0]))
	// 等价于unsafe.Pointer(&b[9])
	fmt.Println(end)
	// 转int类型指针，再取地址
	fmt.Println(*(*int)(end))
	// 9
	p := &Person{age: 30, name: "Bob"}
	//获取到struct s中b字段的地址
	addr := unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Offsetof(p.name))
	//将其转换为一个string的指针，并且打印该指针的对应的值
	fmt.Println(*(*string)(addr))
	// 可以用来进行类型转化，但是不安全
	t := s1{name: "123"}
	j := *(*s2)(unsafe.Pointer(&t))
	fmt.Println(j)
	// 可以通过指针来克隆对象,可以深拷贝
	tp := &t
	t1 := *tp
	fmt.Println("t1 object:", t1)
	t1.id = 500
	fmt.Println("t1 object:", t1)
	fmt.Println("t object:", t)
	dp := &deepObj{stru: t}
	d1 := *dp
	fmt.Println("d1:", d1)
	d1.id = 500
	fmt.Println("d1 object:", d1)
	fmt.Println("d object:", *dp)
	d1.stru.id = 300
	fmt.Println("d1 object:", d1)
	fmt.Println("d object:", *dp)
}
