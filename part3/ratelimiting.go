package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for w := 1; w <= 5; w++ {
		requests <- w
	}
	close(requests)
	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	fmt.Println("End1")
	bustyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		bustyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			bustyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-bustyLimiter
		fmt.Println("request2", req, time.Now())
	}

}
