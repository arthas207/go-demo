package main

import (
	"fmt"
	"go-demo/collection"
	"go-demo/concurrent"
	"go-demo/enums"
	"go-demo/errorhandle"
	demoio "go-demo/io"
	"go-demo/pointer"
	"go-demo/polymorphism"
	"go-demo/reflection"
)

func main() {
	fmt.Println("====collection example start====")
	collection.Example()
	fmt.Println("====concurrent example start====")
	concurrent.Example()
	fmt.Println("====enums example start====")
	enums.Example()
	fmt.Println("====errorhandle example start====")
	errorhandle.Example()
	fmt.Println("====io example start====")
	demoio.Example()
	fmt.Println("====pointer example start====")
	pointer.Example()
	fmt.Println("====polymorphism example start====")
	polymorphism.Example()
	fmt.Println("====reflection example start====")
	reflection.Example()
}
