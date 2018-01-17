package main

import "fmt"

type geometry interface {
	area() int
	perim() int
	setWidth(i int)
}

type rect struct {
	width, heigth int
}

func (r rect) area() int {
	return r.heigth * r.width
}

//use pointer to reflect setting to orig value
func (r *rect) perim() int {
	r.width = 40
	return 2*r.width + 2*r.heigth
}

func (r *rect) setWidth(n int) {
	r.width = n
}

func measure(g geometry) {
	fmt.Println(g.perim())
	fmt.Println(g.area())
	//g.setWidth(30)
	//fmt.Println(g.area())
	//fmt.Println(g.perim())

}

func main() {

	r := rect{10, 20}
	measure(&r)
	fmt.Println(r.heigth)
	fmt.Println(r.width)

}
