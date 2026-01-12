package numberofsteps

import "testing"

func TestNumberOfSteps_Example1(t *testing.T) {
	got := numberOfSteps(14)
	want := 6
	if got != want {
		t.Fatalf("numberOfSteps(14) = %d; want %d", got, want)
	}
}

func TestNumberOfSteps_Example2(t *testing.T) {
	got := numberOfSteps(8)
	want := 4
	if got != want {
		t.Fatalf("numberOfSteps(8) = %d; want %d", got, want)
	}
}

func TestNumberOfSteps_Example3(t *testing.T) {
	got := numberOfSteps(123)
	want := 12
	if got != want {
		t.Fatalf("numberOfSteps(123) = %d; want %d", got, want)
	}
}
