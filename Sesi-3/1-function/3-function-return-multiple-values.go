package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15

	var area, circumference = calculate(diameter)

	fmt.Println("Area: ", area)
	fmt.Println("Circumference: ", circumference)
}

func calculate(d float64) (float64, float64) {
	var area float64 = math.Pi * math.Pow(d/2, 2)

	var circumference = math.Pi * d

	return area, circumference
}
