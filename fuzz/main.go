package main

import "unicode/utf8"

func Reverse(s string) string {
	// FIX: Check if the string is valid UTF-8.
	// If it's not valid, reversing it will corrupt the bytes.
	if !utf8.ValidString(s) {
		return s
	}

	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	input := "The quick brown fox jumps over the lazy dog"
	rev := Reverse(input)
	println("Original:", input)
	println("Reversed:", rev)
	println("Reversed back:", Reverse(rev))
}
