package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 8
	m["k2"] = 123

	fmt.Println("maps:", m)

	v1 := m["k1"]
	m2 := m
	fmt.Println("v1:", v1)

	// change m2 also update m
	m2["k1"] = 12
	fmt.Println("maps:", m)

	fmt.Println("maps2:", m2)
	for k, v := range m2 {
		fmt.Println(k, v)
	}

	m3 := make(map[string]string)
	m3["k1"] = "k1"
	fmt.Println("nil for int type:", m2["k3"], "End")
	fmt.Println("nil for string type:", m3["k3"], "End")

}
