package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("create file ", p)
	f, error := os.Create(p)
	if error != nil {
		panic(error)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

func main() {
	f := createFile("file")
	//executed after the enclosing function
	defer closeFile(f)
	writeFile(f)
}
