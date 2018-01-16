package main

import "fmt"
import "time"

func main() {
	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Work Day")
	}
}
