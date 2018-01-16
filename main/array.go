package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("a:", a)

	fmt.Println("a:", a)
	a[4] = 100
	fmt.Println("a:", a)
}
