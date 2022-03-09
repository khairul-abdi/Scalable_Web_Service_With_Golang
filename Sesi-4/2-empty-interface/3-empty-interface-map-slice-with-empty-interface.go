package main

import "fmt"

func main() {
	rs := []interface{}{1, "Khairul", true, 2, "Diah", true}

	rm := map[string]interface{}{
		"Name":   "Khairul",
		"Status": true,
		"Age":    26,
	}

	fmt.Printf("%+v\n", rs)
	fmt.Printf("%+v\n", rm)
}
