package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Looping Over String (string-by-byte)")
	city := "Jakarta"

	// looping byte/ascii dari tiap karakter
	for i := 0; i < len(city); i++ {
		fmt.Printf("Index: %d, byte: %d, character: %s\n", i, city[i], string(city[i]))
	}

	fmt.Println("---------------------------------------")
	fmt.Println("Looping Over String (byte-by-string)")
	var cityByte = []byte{74, 97, 107, 97, 114, 116, 97}
	fmt.Println(string(cityByte))

	fmt.Println("---------------------------------------")
	fmt.Println("Looping Over String (rune-by-rune)")

	str1 := "makan"
	str2 := "mânca"

	fmt.Printf("str1 byte length => %d\n", len(str1))
	fmt.Printf("str2 byte length => %d\n", len(str2))

	fmt.Printf("str1 character length => %d\n", utf8.RuneCountInString(str1))
	fmt.Printf("str2 character length => %d\n", utf8.RuneCountInString(str2))
}
