package main

import "fmt"

func main() {

	a := make([]string, 3)
	fmt.Println("a:", a)

	a[0] = "100"
	a[1] = "200"
	a[2] = "300"

	b := make([]string, 2)
	copy(b, a)
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	a = append(a, "400")

	fmt.Println(a[:3])
	fmt.Println(a[3:])

	c := a[2:4]
	fmt.Println(c)
}
