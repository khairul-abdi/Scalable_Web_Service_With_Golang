package main

import (
	"fmt"
	"strings"
)

func main() {
	var studenLists = []string{"Khairul", "Diah", "Zainy", "Anggi", "Nia"}
	var find = findStudent(studenLists)

	// fmt.Println(find("khairul", 2))
	res, num := find("khairul", 2)
	fmt.Println(res)
	fmt.Println(num)
}

func findStudent(students []string) func(string, int) (string, int) {
	return func(s string, num int) (string, int) {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}

		if student == "" {
			return fmt.Sprintf("%s doesn't exist!!", s), num
		}
		return fmt.Sprintf("We found %s at position %d dan grade ", s, position+1), num
	}
}
