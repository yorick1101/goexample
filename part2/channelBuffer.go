package main

import "fmt"

func main() {
	// a chanel with value type string
	message := make(chan string, 3)

	message <- "message1"
	message <- "message1"

	fmt.Println(<-message)
	fmt.Println(<-message)

}
