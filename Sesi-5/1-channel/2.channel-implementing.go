package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(3)
	c := make(chan string)

	go introduce("Gojo Satoru", c)
	go introduce("Tanjiro", c)
	go introduce("Minato", c)

	//Menerima data dari channel
	result := <-c
	fmt.Println(result)
	result = <-c
	fmt.Println(result)
	result = <-c
	fmt.Println(result)

	close(c)
}

func introduce(str string, c chan string) {

	result := fmt.Sprintf("Hai, my name is %s ", str)

	// Mengirim data kepada channel
	c <- result
}
