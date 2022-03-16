package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonString = `
	{
		"full_name": "Khairul",
		"email": "kabdi384@gmail.com",
		"age": 26
	}`

	var result map[string]interface{}

	var err = json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("full_name:", result["full_name"])
	fmt.Println("email:", result["email"])
	fmt.Println("age:", result["age"])
}
