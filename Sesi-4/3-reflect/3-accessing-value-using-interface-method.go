package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	var nilai = reflectValue.Interface().(int)
	fmt.Println(nilai)
	fmt.Println(reflect.TypeOf(nilai))
}
