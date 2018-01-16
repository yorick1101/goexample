package main

import "fmt"

func plus(a int, b int) (int, int) {
	return a, b
}

func main() {
	res, res2 := plus(1, 2)
	fmt.Printf("%d, %d\n", res, res2)

	res3, _ := plus(1, 2)

	fmt.Printf("%d\n", res3)

}
