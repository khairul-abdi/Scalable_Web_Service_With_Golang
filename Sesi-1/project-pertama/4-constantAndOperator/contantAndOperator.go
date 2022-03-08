package main

import "fmt"

func main() {
	fmt.Println("CONSTANT")
	// const full_name string //ERROR kalau tidak langsung di isi valuenya

	const fullname string = "Khairul"
	fmt.Printf("Hello %s", fullname)

	fmt.Println("--------------------------------------------------")
	fmt.Println("Operator")
	var value = 2 + 2*3
	fmt.Println("Tanpa ada prioritas => ", value)
	value = (2 + 2) * 3
	fmt.Println("Prioritas dengan diberi tanda () => ", value)

	fmt.Println("--------------------------------------------------")
	fmt.Println("Operators (Relational Operators)")

	var firstCondition bool = 2 < 3
	var secondCondition bool = "khairul" == "Khairul"
	var thirdCondition bool = 10 != 2.3
	var fourthCondition bool = 11 <= 11

	fmt.Println("first condition: ", firstCondition)
	fmt.Println("second condition: ", secondCondition)
	fmt.Println("third condition: ", thirdCondition)
	fmt.Println("fourth condition: ", fourthCondition)

	fmt.Println("--------------------------------------------------")
	fmt.Println("Operators (Logical Operators)")

	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	fmt.Printf("wrong && right \t(%t)\n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("wrong || right \t(%t)\n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!wrong \t\t(%t)\n", wrongReverse)

}
