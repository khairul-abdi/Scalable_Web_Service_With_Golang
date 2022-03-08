package main

import "fmt"

func main() {
	numberList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := sum(numberList...)
	fmt.Printf("Result: %d", result)
}

func sum(numList ...int) int {
	var result int

	for _, v := range numList {
		result += v
	}

	return result
}
