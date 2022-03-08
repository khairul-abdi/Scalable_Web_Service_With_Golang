package main

import "fmt"

func main() {
	var firstNumber int = 4
	var secondNumber *int = &firstNumber

	fmt.Println("firstNumber (value) \t\t:", firstNumber)
	fmt.Println("firstNumber (memori address) \t:", &firstNumber)

	fmt.Println("secondNumber (value) \t\t:", *secondNumber)
	fmt.Println("secondNumber (memori address) \t:", secondNumber)
}
