package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	var employee = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "Khairul", age: 26}, division: "Backend Developer"},
		{person: Person{name: "Diah", age: 21}, division: "Data Analyst"},
		{person: Person{name: "Zainy", age: 24}, division: "Backend Developer"},
		{person: Person{name: "Anggi", age: 19}, division: "Devops"},
		{person: Person{name: "Nia", age: 18}, division: "CEO"},
	}

	for _, v := range employee {
		fmt.Printf("%+v\n", v)
	}
}
