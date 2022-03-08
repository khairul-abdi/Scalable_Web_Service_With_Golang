package main

import "fmt"

func main() {
	greet("Khairul", "Jalan Tanah Kusir")
}

func greet(name, address string) {
	fmt.Println("Hello there! My name is", name)
	fmt.Println("I live in", address)
}
