package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan map[string]interface{}, 2)

	go func(c chan map[string]interface{}) {
		for i := 0; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
			res := map[string]interface{}{"name": "khairul"}
			c <- res
			fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
		}

		close(c)
	}(c1)

	fmt.Println("main goroutine sleeps 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c1 { //v = <- c1
		fmt.Println("main goroutine received value from channel:", v)
	}

}
