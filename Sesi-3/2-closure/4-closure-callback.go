package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = []string{"Khairul", "Diah", "Zainy", "Anggi", "Nia"}
	var dataContainI = filter(data, func(each string) bool {
		return strings.Contains(each, "i")
	})

	var dataLength5 = filter(data, func(each string) bool {
		return len(each) >= 5
	})

	fmt.Printf("Data Asli: %#v\n", data)
	fmt.Printf("Filter ada huruf \"i\": %#v\n", dataContainI)
	fmt.Printf("Filter panjang karakter lebih dari 5 : %#v\n", dataLength5)
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}

	return result
}
