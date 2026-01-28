package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Printf("Index: %v, Arg: %v\n", i, arg)
	}
	fmt.Println(s)
}
