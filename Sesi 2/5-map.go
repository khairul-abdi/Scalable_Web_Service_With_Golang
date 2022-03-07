package main

import (
	"fmt"
	"strings"
)

func main() {

	var person = map[string]string{}
	person["name"] = "Khairul"
	person["age"] = "26"
	person["address"] = "Jalan Tanah Kusir"

	fmt.Println("name:", person["name"])
	fmt.Println("age:", person["age"])
	fmt.Println("address:", person["address"])

	fmt.Println(strings.Repeat("#", 20))
	var person2 = map[string]string{
		"name":    "Khairul",
		"age":     "26",
		"address": "Jalan Tanah Kusir",
	}

	fmt.Printf("%#v\n", person2)
	fmt.Println(strings.Repeat("#", 20))

	fmt.Println("----------------------------")
	fmt.Println("Map (Looping with map)")
	for key, value := range person2 {
		fmt.Println(key, ":", value)
	}

	fmt.Println("----------------------------")
	fmt.Println("Map (Deleting value 1)")

	fmt.Println("Before deleting: ", person2)
	delete(person2, "age")
	fmt.Println("After deleted: ", person2)

	fmt.Println("----------------------------")
	fmt.Println("Map (Deleting value 2)")

	value, exist := person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value doesn`t exist")
	}

	delete(person, "age")

	value, exist = person["age"]
	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value has been deleted")
	}

	fmt.Println("----------------------------")
	fmt.Println("Map (Combining slice with map)")
	var people = []map[string]string{
		{"name": "Khairul", "age": "26"},
		{"name": "Zainy", "age": "24"},
		{"name": "Anggi", "age": "20"},
		{"name": "Nia", "age": "19"},
	}

	value, exist = people[0]["age"]
	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value doesn`t exist")
	}

	for i, person := range people {
		fmt.Printf("Index: %d, name: %s, age: %s\n", i, person["name"], person["age"])
	}
}
