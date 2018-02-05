package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var p = fmt.Println

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("dat1")
	check(err)
	fmt.Print(string(dat))

	f, err := os.Open("dat1")
	check(err)

	b1 := make([]byte, 10)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	//second paramete:r 0 means relative to the origin of the file, 1 means relative to the current offset, and 2 means relative to the end
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	//read at least 2 bytes into b3
	//Error returns is reading fewer bytes

	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
	defer f.Close()
}
