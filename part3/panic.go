package main

import "os"

func main() {
	panic("a problem")

	_, error := os.Create("/temp/file")
	if error != nil {
		panic(error)
	}
}
