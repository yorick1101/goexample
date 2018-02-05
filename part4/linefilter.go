package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//The returned value of type *os.File implements the io.Reader interface.
	file, _ := os.Open("test.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println("L:", ucl)

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
