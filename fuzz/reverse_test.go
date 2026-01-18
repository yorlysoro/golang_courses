package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"The quick brown fox", "xof nworb kciuq ehT"},
		{"こんにちは", "はちにんこ"},
	}

	for _, c := range cases {
		got := Reverse(c.input)
		if got != c.expected {
			t.Errorf("Reverse(%q) == %q, want %q", c.input, got, c.expected)
		}
	}
}

func FuzzReverse(f *testing.F) {
	testCases := []string{"hello", "world", "The quick brown fox", "こんにちは"}
	for _, tc := range testCases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		// If the input isn't valid UTF-8, our Reverse function
		// (and the concept of runes) doesn't apply normally.
		if !utf8.ValidString(orig) {
			return
		}
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
