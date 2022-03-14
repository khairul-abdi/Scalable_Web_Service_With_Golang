package main

import "fmt"

func main() {
	go goroutine()
}

func goroutine() {
	fmt.Println("Hello")
}
