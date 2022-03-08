package main

import (
	"fmt"
	"strings"
)

func main() {
	var fruits = []string{"apple", "banana", "manggo"}
	_ = fruits

	fmt.Println("----------------------------")
	fmt.Println("Slice (make function)")
	var fruits2 = make([]string, 3)
	_ = fruits2
	fmt.Printf("%#v", fruits2)

	fmt.Println("----------------------------")
	fmt.Println("Slice (append function)")

	var fruits3 = make([]string, 3)
	_ = fruits3
	fruits3[0] = "apple"
	fruits3[1] = "banana"
	fruits3[2] = "manggo"

	fmt.Printf("%#v\n", fruits3)

	var fruits4 = make([]string, 0)
	fruits4 = append(fruits4, "durian")
	fruits4 = append(fruits4, "apple", "banana", "manggo")
	fmt.Printf("%#v\n", fruits4)

	fmt.Println("----------------------------")
	fmt.Println("Slice (append function with ellipsis)")
	var fruits5 = []string{"apple", "banana", "manggo"}
	var fruits6 = []string{"durian", "pineapple", "starfruit"}

	fruits5 = append(fruits5, fruits6...)
	fmt.Printf("%#v", fruits5)

	fmt.Println("----------------------------")
	fmt.Println("Slice (copy function)")

	var fruits7 = []string{"apple", "banana", "mango", "watermelon"}
	var fruits8 = []string{"durian", "pineapple", "starfruit"}

	nn := copy(fruits7, fruits8)
	fmt.Println("Fruits 7 => ", fruits7)
	fmt.Println("Fruits 8 => ", fruits8)
	fmt.Println("Copied elements => ", nn)

	fmt.Println("----------------------------")
	fmt.Println("Slice (Slicing)")
	var fruits9 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	var fruits10 = fruits9[1:4]
	fmt.Printf("%#v\n", fruits10)

	var fruits11 = fruits9[0:]
	fmt.Printf("%#v\n", fruits11)

	var fruits12 = fruits9[:3]
	fmt.Printf("%#v\n", fruits12)

	var fruits13 = fruits9[:]
	fmt.Printf("%#v\n", fruits13)

	fmt.Println("----------------------------")
	fmt.Println("Slice (Combining slicing and append)")

	var fruits14 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	fruits14 = append(fruits14[:3], "rambutan")
	fmt.Printf("%#v", fruits14)

	fmt.Println("----------------------------")
	fmt.Println("Slice (Backing array)")

	var fruits15 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	var fruits16 = fruits15[2:4]
	fruits16[0] = "rambutan"

	fmt.Println("Fruits14 => ", fruits15)
	fmt.Println("Fruits15 => ", fruits16)

	fmt.Println("----------------------------")
	fmt.Println("Slice (Cap function)")

	var fruits17 = []string{"apple", "mango", "durian", "banana"}

	fmt.Println("Fruits1 cap => ", cap(fruits17))
	fmt.Println("Fruits1 len => ", len(fruits17))
	fmt.Println(strings.Repeat("#", 20))

	var fruits18 = fruits17[0:3]

	fmt.Println("Fruits1 cap => ", cap(fruits18))
	fmt.Println("Fruits1 len => ", len(fruits18))
	fmt.Println(strings.Repeat("#", 20))

	var fruits19 = fruits17[1:]
	fmt.Println("Fruits1 cap => ", cap(fruits19))
	fmt.Println("Fruits1 len => ", len(fruits19))

	fmt.Println("----------------------------")
	fmt.Println("Slice (Creating a new backing array)")
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println("cars: ", cars)
	fmt.Println("NewCars: ", newCars)
}
