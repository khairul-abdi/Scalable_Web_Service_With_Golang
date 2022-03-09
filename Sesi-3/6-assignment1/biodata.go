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
	{Id: 2, Name: "Evrin", Address: "Jalan Rasuna", Profession: "Backend Developer", Reason: "Belajar programming"},
	{Id: 3, Name: "Dedy Chandra", Address: "Jalan Sudirman", Profession: "Backend Developer", Reason: "Belajar bahasa go"},
	{Id: 4, Name: "Ghifari", Address: "Jalan Tanah Kusir", Profession: "Backend Developer", Reason: "Belajar bahasa baru"},
	{Id: 5, Name: "Alvin", Address: "Jalan Tanah Kusir", Profession: "Backend Developer", Reason: "Belajar bahasa go"},
	{Id: 6, Name: "Faikar", Address: "Jalan Tanah Kusir", Profession: "Backend Developer", Reason: "Belajar bahasa go milik google"},
}

func main() {

	argument := os.Args[1]
	arg, err := strconv.Atoi(argument)
	if err != nil {
		fmt.Printf("Not Convert to int => Type Data: %T, Value: %v\n", arg, arg)
	}

	result := filter(students, arg)
	resJson, _ := json.Marshal(result)
	fmt.Println(string(resJson))
}

func filter(students []helpers.Student, arg int) helpers.Student {
	var result helpers.Student
	for _, v := range students {
		if v.Id == arg {
			result = v
		}
	}

	return result
}
