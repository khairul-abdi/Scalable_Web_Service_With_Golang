package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loopings (first way of looping)")
	for i := 0; i < 3; i++ {
		fmt.Printf("Angka: %d\n", i)
	}

	fmt.Println("-----------------------------------------------")
	fmt.Println("Loopings (second way of looping)")

	i := 0

	for i < 3 {
		fmt.Printf("Index ke %d\n", i)
		i++
	}

	fmt.Println("-----------------------------------------------")
	fmt.Println("Loopings (third way of looping)")

	i = 0
	for {
		fmt.Printf("Index ke %d\n", i)

		i++
		if i == 3 {
			break
		}
	}

	fmt.Println("-----------------------------------------------")
	fmt.Println("Loopings (break and continue keywords)")

	for i := 0; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
		if i > 6 {
			break
		}
		fmt.Printf("Index ke %d\n", i)
	}

	fmt.Println("-----------------------------------------------")
	fmt.Println("Loopings (Nested Looping)")
	for i := 0; i <= 4; i++ {
		for j := i; j <= 4; j++ {
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}

	fmt.Println("-----------------------------------------------")
	fmt.Println("Loopings (Label)")

outerLoop:
	for i := 0; i < 5; i++ {
		fmt.Printf("Index ke - %d\n", i)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}

}
