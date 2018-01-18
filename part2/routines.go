package main

import "fmt"

func f(from string) {
	for i := 0; i < 20; i++ {
		fmt.Println("from:", from, i)
	}
}

func main() {
	go f("async")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	f("sync")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
