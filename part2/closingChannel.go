package main

import (
	"fmt"
)

func main() {
	// a chanel with value type string
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			//the second parameter indicates the channel is closed and no values exist in the channel
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- false
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	//block the main to wait for done message
	<-done
}
