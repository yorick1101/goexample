package main

import (
	"fmt"
)

func main() {
	// a chanel with value type string
	message := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message yet")
	}

	msg := "hi"
	select {
	case message <- msg:
		fmt.Println("send message", msg)
	default:
		fmt.Println("message not sent")
	}

	select {
	case msg := <-message:
		fmt.Println("Received Message:", msg)
	case sig := <-signals:
		fmt.Println("Received Signal", sig)
	default:
		fmt.Println("no activity")
	}

}
