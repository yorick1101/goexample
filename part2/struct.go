package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "BoD", age: 2})

	p1 := person{}
	p1.name = "Hau"
	p1.age = 30
	fmt.Println(p1)

	p2 := &p1
	p2.name = "Hau1"
	p2.age = 31
	fmt.Println(p1)
	fmt.Println(p2)
}
