package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("Type variabel :", reflectValue.Type())
	fmt.Println("Nilai variabel:", reflectValue.Interface())
}
