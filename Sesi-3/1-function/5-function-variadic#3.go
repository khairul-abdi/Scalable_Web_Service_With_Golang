package main

import (
	"fmt"
	"strings"
)

func main() {
	profile("Khairul", "ikan bakar", "ketoprak", "siomay")
}

func profile(name string, favFoods ...string) {

	mergeFavFoods := strings.Join(favFoods, ", ")
	fmt.Println("Hello there!! I`m", name)
	fmt.Println("I really love to eat", mergeFavFoods)
}
