package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Use strings.Join to concatenate command-line arguments with spaces
	// and compare it to os.Args directly.
	// Calculate the time taken to execute both methods.
	// Use os.Args[1:] to skip the program name.
	fmt.Println(os.Args[1:])
	fmt.Println("Using strings.Join:")
	fmt.Println(strings.Join(os.Args[1:], " "))
}
