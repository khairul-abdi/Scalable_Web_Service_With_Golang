package main

import "fmt"

func main() {
	fmt.Println("Variable with data type")
	var name string = "Khairul"
	var age int = 26

	//Re asign
	name = "Abdi"
	age = 25

	//Error karna tipe data tidak sesuai dengan value yang diisi
	// name = 26
	// age = "Abdi"

	fmt.Println("Ini adalaha namanya ==> ", name)
	fmt.Println("Ini adalaha umurnya ==> ", age)

	fmt.Println("----------------------------------------------")
	fmt.Println("Variable without data type")

	var name2 = "Khairul"
	var age2 = 26
	fmt.Printf("%T, %T\n", name2, age2)

	fmt.Println("----------------------------------------------")
	fmt.Println("Variable without data type(Short Declaration)")

	name3 := "Khairul"
	age3 := 26
	fmt.Printf("%T, %T\n", name3, age3)

	fmt.Println("----------------------------------------------")
	fmt.Println("Multiple variable declarations")

	var student1, student2, student3 string = "satu", "dua", "tiga"
	var first, two, three int
	first, two, three = 1, 2, 3

	fmt.Println(student1, student2, student3)
	fmt.Println(first, two, three)

	fmt.Println("----------------------------------------------")
	fmt.Println("Multiple variable declarations with type different")

	var name4, age4, address = "Khairul", 26, "jl.Tanah kusir"
	first2, second2, third2 := "1", 2, true

	fmt.Println(name4, age4, address)
	fmt.Println(first2, second2, third2)

	fmt.Println("----------------------------------------------")
	fmt.Println("Underscore variable")
	var firstVariable string
	var name5, age5, address5 = "Khairul", 26, "Jl. Tanah Kusir"

	_, _, _, _ = firstVariable, name5, age5, address5

	fmt.Println("----------------------------------------------")
	fmt.Println("fmt.Printf Usage")

	first3, second3 := 1, "2"
	fmt.Printf("Tipe data variable first adalah %T \n", first3)
	fmt.Printf("Tipe data variable first adalah %T \n", second3)

	var nama = "Khairul"
	var umur = 26
	var alamat = "Jalan Tanah Kusir"

	fmt.Printf("Halo namaku %s, umur aku adalah %d, dan aku tinggal di %s", nama, umur, alamat)
}
