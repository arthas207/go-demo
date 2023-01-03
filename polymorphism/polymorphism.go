package polymorphism

import (
	"fmt"
)

type Car interface {
	drive()
}

type BMW struct{}

func (Bmw BMW) drive() {
	fmt.Println("bmw")
}

type LAM struct{}

func (Lam LAM) drive() {
	fmt.Println("lam")
}

type TOYOTA struct{}

func (Toyota TOYOTA) drive() {
	fmt.Println("toyota")
}

func Example() {
	var car Car
	car = BMW{}
	car.drive()
	car = LAM{}
	car.drive()
	car = TOYOTA{}
	car.drive()
}
