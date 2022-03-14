package main

import "fmt"

func main() {
	c := make(chan string)

	go print("Gojo Satoru", c)

	//Menerima data dari channel
	result := <-c
	fmt.Println(result)
}

func print(str string, c chan string) {

	result := fmt.Sprintf("Hai, my name is %s ", str)

	// Mengirim data kepada channel
	c <- result
}
