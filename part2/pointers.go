package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroval2(ival *int) {
	*ival = 0
}

func main() {

	i := 1
	fmt.Println(i)

	zeroval(i)
	fmt.Println(i)
	zeroval2(&i)
	fmt.Println(i)
}
