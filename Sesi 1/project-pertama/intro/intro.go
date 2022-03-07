package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello", "World")
	fmt.Println("Hello", " ", "World")

	var name = "KHAIRUL"
	var age = 26
	fmt.Printf("%T, %T\n", name, age)
	fmt.Println(reflect.TypeOf(name), reflect.TypeOf(age))
}
