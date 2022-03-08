package main

import "fmt"

func main() {
	var currentYear = 2022

	if age := currentYear - 1995; age < 17 {
		fmt.Println("Kamu belum boleh membuat kartu sim")
	} else {
		fmt.Println("Kamu sudah boleh membuat kartu sim")
	}

	fmt.Println("--------------------------------------------------")
	fmt.Println("Bagian Switch")

	score := 8

	switch score {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	fmt.Println("--------------------------------------------------")
	fmt.Println("Switch with relational operators")

	score2 := 2

	switch {
	case score2 == 8:
		fmt.Println("perfect")
	case (score2 < 8) && (score2 > 3):
		fmt.Println("not bad")
	default:
		{
			fmt.Println("Study harder")
			fmt.Println("you need to learn more")
		}
	}

	fmt.Println("--------------------------------------------------")
	fmt.Println("Switch (fallthrough keyword)")
	var score3 int = 6

	switch {
	case score3 == 8:
		fmt.Println("perfect")
		fallthrough
	case (score3 < 8) && (score3 > 3):
		fmt.Println("not bad")
		fallthrough
	case score3 < 5:
		fmt.Println("It is ok, but please study harder")
		fallthrough
	default:
		{
			fmt.Println("study harder")
			fmt.Println("You don`t have a good score yet")
		}
	}

	fmt.Println("--------------------------------------------------")
	fmt.Println("Nested Conditions")

	score4 := 0

	if score4 > 7 {
		switch score4 {
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice!")
		}
	} else {
		if score4 == 5 {
			fmt.Println("not bad")
		} else if score4 == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if score4 == 0 {
				fmt.Println("try harder!")
			}
		}
	}

}
