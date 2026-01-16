package main

import "fmt"

func arrays1() {
	// 1. Declaration without initialization
	// This creates an array of 5 integers, all defaulted to 0.
	var a [5]int
	fmt.Println("Empty array:", a)

	// 2. Declaration with initialization
	// We define the size and the values immediately.
	b := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Initialized array:", b)

	// 3. Let Go count the elements
	// Using '...' allows the compiler to determine the length based on the elements.
	c := [...]string{"Go", "Python", "Java"}
	fmt.Println("Compiler-counted array:", c)
	fmt.Println("Length of c:", len(c))
}

func arrays2() {
	nums := [3]int{10, 20, 30}

	// Setting a value at a specific index
	nums[1] = 50

	// Getting a value
	fmt.Println("Value at index 1:", nums[1])

	// Iterating using a standard for loop
	fmt.Println("Iterating with for loop:")
	for i := 0; i < len(nums); i++ {
		fmt.Printf("Index %d: %d\n", i, nums[i])
	}

	// Iterating using range (The idiomatic Go way)
	fmt.Println("Iterating with range:")
	for index, value := range nums {
		fmt.Printf("Index %d has value %d\n", index, value)
	}
}

// DVD defines the properties of a movie disk
type DVD struct {
	Name        string
	ReleaseYear int
	Director    string
}

// Showable is the interface that requires a String() method
// This allows any type that implements String() to be treated as Showable
type Showable interface {
	String() string
}

// String implements the logic to format the DVD data nicely.
// This method makes the DVD struct satisfy the Showable interface.
func (d DVD) String() string {
	return fmt.Sprintf("🎬 %-20s | 👤 Dir: %-15s | 📅 %d", d.Name, d.Director, d.ReleaseYear)
}

func DVDs() {
	// Creating a collection of 15 DVDs using an array
	// Syntax: [size]Type
	var collection [15]DVD

	// Manually populating some data for demonstration
	collection[0] = DVD{"Inception", 2010, "Christopher Nolan"}
	collection[1] = DVD{"The Matrix", 1999, "The Wachowskis"}
	collection[2] = DVD{"Pulp Fiction", 1994, "Quentin Tarantino"}
	collection[3] = DVD{"Interstellar", 2014, "Christopher Nolan"}
	collection[4] = DVD{"Parasite", 2019, "Bong Joon-ho"}

	// Filling the rest with placeholders to reach 15
	for i := 5; i < 15; i++ {
		collection[i] = DVD{Name: fmt.Sprintf("Movie Vol. %d", i+1), ReleaseYear: 2000 + i, Director: "Unknown"}
	}

	fmt.Println("--- My DVD Collection ---")

	// Iterating through the array and calling the interface method
	for _, item := range collection {
		// Even though 'item' is a DVD struct, it satisfies the Showable interface
		displayInfo(item)
	}
}

// displayInfo takes the Showable interface as a parameter
func displayInfo(s Showable) {
	fmt.Println(s.String())
}
