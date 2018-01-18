package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	// a chanel with value type string
	done := make(chan bool, 1)

	go worker(done)

	fmt.Println("wait")
	<-done
	fmt.Println("receive done")

}
