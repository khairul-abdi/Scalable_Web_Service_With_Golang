package main

import "fmt"

func main() {
	c := make(chan string)

	students := []string{"Khairul", "Zainy", "Anggi", "Nia"}

	for _, v := range students {
		go func(student string) {
			fmt.Println("Student", student)
			result := fmt.Sprintf("Hai, my name is %s", student)
			c <- result
		}(v)
	}

	for i := 0; i < len(students); i++ {
		print(c)
	}

	close(c)
}

func print(c chan string) {
	fmt.Println(<-c)
}
