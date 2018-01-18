package main

import "fmt"

func main() {
	// a chanel with value type string
	message := make(chan string)

	go func() {
		for true {
			msg := <-message
			fmt.Println(msg)
		}
	}()

	go func() {
		i := 0
		// compile failed for i++ < 100
		for i < 100 {
			message <- fmt.Sprintf("ping %d", i)
			i++
		}
	}()

	var done string
	fmt.Scanln(&done)
}
