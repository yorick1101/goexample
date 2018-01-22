package main

import (
	"fmt"
	"time"
)

func main() {
	// a chanel with value type string
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "one"
	}()

	select {
	case msg := <-c1:
		fmt.Println("received from c1:", msg)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	select {
	case msg := <-c2:
		fmt.Println("Received from c2:", msg)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}

}
