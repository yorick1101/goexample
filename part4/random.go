package main

import (
	"fmt"
	"math/rand"
	"time"
)

var p = fmt.Println

func main() {

	p(rand.Intn(100), ",")
	p(rand.Intn(100))
	p(" ")

	p(rand.Float64())
	p(rand.Float64()*5 + 5)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	//If you seed a source with the same number, it produces the same sequence of random numbers.
	p(r1.Intn(100), ",")
	p(r1.Intn(100))
	p(" ")

}
