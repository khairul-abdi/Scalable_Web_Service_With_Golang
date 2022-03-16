package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func main() {
	var jsonString = `[
		{
			"full_name": "Khairul",
			"email": "kabdi384@gmail.com",
			"age": 26
		},
		{
			"full_name": "Diah",
			"email": "diah@gmail.com",
			"age": 22
		}
	]`

	var result []Employee

	err := json.Unmarshal([]byte(jsonString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, v := range result {
		fmt.Printf("Index %d: %+v \n", i, v)
	}
}
