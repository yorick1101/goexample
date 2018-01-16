package main

import "fmt"

func intSeq() func() int {
	i := 0
	//anynymously
	return func() int {
		i += 1
		return i
	}
}

func main() {
	next := intSeq()

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

}
