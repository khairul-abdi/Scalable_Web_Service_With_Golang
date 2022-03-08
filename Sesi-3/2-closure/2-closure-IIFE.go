package main

import "fmt"

func main() {

	var isPalindrome = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}

		return str == temp
	}("katak")

	fmt.Println(isPalindrome)
}
