package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstPerson string = "John"
	var secondPerson *string = &firstPerson

	fmt.Println("firstPerson (value) :", firstPerson)
	fmt.Println("firstPerson (memori address) : ", &firstPerson)
	fmt.Println("secondNumber (value) : ", *secondPerson)
	fmt.Println("secondNumber (memoru address) : ", secondPerson)

	fmt.Println("\n", strings.Repeat("#", 30), "\n")

	*secondPerson = "Doe"

	fmt.Println("firstPerson (value) :", firstPerson)
	fmt.Println("firstPerson (memori address) : ", &firstPerson)
	fmt.Println("secondNumber (value) : ", *secondPerson)
	fmt.Println("secondNumber (memoru address) : ", secondPerson)
}
