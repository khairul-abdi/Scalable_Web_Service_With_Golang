package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	var s1 = &student{Name: "john wick", Grade: 2}
	fmt.Println("name :", s1.Name)

	var reflecValue = reflect.ValueOf(s1)
	var method = reflecValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println("nama :", s1.Name)

}
