package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var people = []Person{
		{name: "Khairul", age: 26},
		{name: "Diah", age: 21},
		{name: "Zainy", age: 24},
		{name: "Anggi", age: 19},
		{name: "Nia", age: 18},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}
}
