package greetings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	name := "Yorlys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Yorlys") = %q, %v, want match for %#q, <nil>`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	_, err := Hello("")
	if err == nil {
		t.Fatalf(`Hello("") = _, %v, want _, error`, err)
	}
}
