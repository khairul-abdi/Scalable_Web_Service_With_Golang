package main

import "fmt"

func main() {
	studentList := print("Khairul", "Diah", "Zainy", "Anggi", "Nia")

	fmt.Println("%#v", studentList)
}

func print(name ...string) []map[string]string {
	var result []map[string]string

	for i, v := range name {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}

	return result
}
