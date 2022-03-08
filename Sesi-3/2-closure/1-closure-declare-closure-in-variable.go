package main

import "fmt"

func main() {
	var evenNumbers = func(numbers ...int) []int {
		var result []int

		for i := 0; i < len(numbers); i++ {
			if numbers[i]%2 == 0 {
				result = append(result, numbers[i])
			}
		}
		return result
	}

	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 11, 12, 14, 15, 28}
	fmt.Println(evenNumbers(numbers...))
}
