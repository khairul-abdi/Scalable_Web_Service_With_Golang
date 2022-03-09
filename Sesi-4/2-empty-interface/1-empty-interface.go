package main

import "fmt"

func main() {
	var randomValue interface{}

	_ = randomValue

	randomValue = "Jalan Tanah Kusir"
	randomValue = 20
	randomValue = true
	randomValue = []string{"Khairul", "Diah"}

	fmt.Println(randomValue)
}
