package main

import (
	"flag"
	"fmt"
)

func main() {

	wordPtr := flag.String("word", "foo", "a string")
	numPtr := flag.Int("numb", 43, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()

	//display the flags have been set
	fmt.Println(flag.NFlag())
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
	//number of the arguments besides flags
	fmt.Println(flag.NArg())

}
