package errorhandle

import "fmt"

func Example() {
	go func(age int) {
		defer func() {
			fmt.Println("recover1 start")
			recover()
		}()
		if age > 90 {
			panic("test panic")
		}
		fmt.Println("no panic")
	}(100)
	defer func() {
		fmt.Println("recover2 start")
		recover()
	}()
	func() {
		panic("a problem occur")
	}()
	fmt.Println("After panic()")
}
