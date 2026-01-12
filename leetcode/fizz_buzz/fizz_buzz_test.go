package fizzbuzz

import "testing"

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestFizzBuzzExample1(t *testing.T) {
	n := 3
	want := []string{"1", "2", "Fizz"}
	if got := fizzBuzz(n); !equalSlices(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestFizzBuzzExample2(t *testing.T) {
	n := 5
	want := []string{"1", "2", "Fizz", "4", "Buzz"}
	if got := fizzBuzz(n); !equalSlices(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestFizzBuzzExample3(t *testing.T) {
	n := 15
	want := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	if got := fizzBuzz(n); !equalSlices(got, want) {
		t.Fatalf("got %v want %v", got, want)
	}
}

/*func TestFizzBuzzLen(t *testing.T) {
	n := 10 ^ 4
	got := fizzBuzz(n)
	if len(got) != n {
		t.Fatalf("got length %d want %d", len(got), n)
	}
}*/
