package main

import (
	"fmt"
)

func sumInts(m map[string]int) int {
	total := 0
	for _, n := range m {
		total += n
	}
	return total
}

func sumFloats(m map[string]float64) float64 {
	total := 0.0
	for _, n := range m {
		total += n
	}
	return total
}

func sumIntsOrFloats[K comparable, V int | float64](m map[K]V) V {
	var total V
	for _, n := range m {
		total += n
	}
	return total
}

type Number interface {
	int | float64
}

func sumNumbers[K comparable, V Number](m map[K]V) V {
	var total V
	for _, n := range m {
		total += n
	}
	return total
}

func main() {
	ints := map[string]int{"first": 1, "second": 2, "third": 3}
	floats := map[string]float64{"first": 1.1, "second": 2.2, "third": 3.3}

	println("Non -generic sums:")
	println("Sum of ints:", sumInts(ints))
	println("Sum of floats:", sumFloats(floats))
	fmt.Printf("\nGeneric sums: %v and %v\n",
		sumIntsOrFloats(map[string]int{"first": 1, "second": 2, "third": 3}),
		sumIntsOrFloats(map[string]float64{"first": 1.1, "second": 2.2, "third": 3.3}))
	fmt.Printf("\nGeneric sums without specifying inferred types: %v and %v\n",
		sumIntsOrFloats(ints),
		sumIntsOrFloats(floats))
	fmt.Printf("\nGeneric sums using Number interface: %v and %v\n",
		sumNumbers(ints),
		sumNumbers(floats))
}
