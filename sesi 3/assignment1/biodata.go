package main

import (
	"assignment1/helpers"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

var students = []helpers.Student{
	{Id: 1, Name: "Khairul", Address: "Jalan Tanah Kusir", Profession: "Backend Developer", Reason: "Belajar bahasa baru"},
	{Id: 2, Name: "Diah", Address: "Jalan Tanah Kusir", Profession: "Backend Developer", Reason: "Belajar bahasa baru"},
}

func main() {

	Argument := os.Args[1]
	Arg, err := strconv.Atoi(Argument)
	if err != nil {
		fmt.Printf("Not Convert to int => Type Data: %T, Value: %v\n", Arg, Arg)
	}

	result := filter(students)
	resJson, _ := json.Marshal(result(Arg))
	fmt.Println(string(resJson))
}

func filter(students []helpers.Student) func(int) helpers.Student {

	return func(arg int) helpers.Student {
		var result helpers.Student
		for _, v := range students {
			if v.Id == arg {
				result = v
			}
		}

		return result
	}
}
