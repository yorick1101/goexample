package main

import (
	"fmt"
	"time"
)

func main() {
	// a chanel with value type string
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-c1:
			fmt.Println("Received from c1:", msg)
		case msg := <-c2:
			fmt.Println("Received from c2:", msg)
		}

	}
}
