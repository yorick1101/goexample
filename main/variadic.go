package main

import "fmt"

func sum(nums ...int) int {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
	return total
}

func main() {
	sum(1, 3, 4, 5)

}
