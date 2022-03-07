package main

import "fmt"

func main() {
	fmt.Println("----------------------------------------------")
	fmt.Println("Number (integer)")

	// var first = 89
	// var second = -12

	// SEBAIKNYA KITA TENTUKAN TIPE DATA AGAR TIDAK BOROS MEMORI
	var first uint8 = 89
	var second int8 = -12

	fmt.Printf("Tipe data first: %T\n", first)
	fmt.Printf("Tipe data second: %T\n", second)

	fmt.Println("----------------------------------------------")
	fmt.Println("Number (float)")

	var decimalNumber float32 = 3.63
	fmt.Printf("Desimal Number: %f\n", decimalNumber)
	fmt.Printf("Desimal Number: %.3f\n", decimalNumber)

	fmt.Println("----------------------------------------------")
	fmt.Println("Number (float)")

	var condition bool = true
	fmt.Printf("is it permitted? %t \n", condition)

	fmt.Println("----------------------------------------------")
	fmt.Println("String")
	var message string = "Halo"
	fmt.Printf("Pesan %s", message)
}
