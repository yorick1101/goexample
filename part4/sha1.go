package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

var p = fmt.Println

func main() {

	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	p(s)
	//print as hex string
	fmt.Printf("%x\n", bs)
	//md5
	m := md5.New()
	m.Write([]byte(s))
	ms := m.Sum(nil)
	fmt.Printf("%x\n", ms)

}
