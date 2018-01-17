package main

import "fmt"

type rect struct {
	width, heigth int
}

func (r *rect) area() int {
	return r.heigth * r.width
}

func (r *rect) perim() int {
	return 2*r.width + 2*r.heigth
}

func main() {

	r := rect{10, 20}
	fmt.Println(r.area())
	fmt.Println(r.perim())
}
