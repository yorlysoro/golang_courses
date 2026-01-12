package runningsum1darray

import "testing"

func TestRunningSumEmpty(t *testing.T) {
	_, err := runningsum([]int{})
	if err == nil {
		t.Fatalf(`runningsum([]int{}) = _, %v, want _, error`, err)
	}
}

func TestRunningSumExceedsMaxLength(t *testing.T) {
	nums := make([]int, 1001)
	_, err := runningsum(nums)
	if err == nil {
		t.Fatalf(`runningsum(large slice) = _, %v, want _, error`, err)
	}
}

func TestRunningSumElementOutOfRange(t *testing.T) {
	nums := []int{10, 20, 30, 1000000}
	_, err := runningsum(nums)
	if err == nil {
		t.Fatalf(`runningsum(%v) = _, %v, want _, error`, nums, err)
	}
}
func TestRunningSumExample1(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	want := []int{1, 3, 6, 10}
	got, err := runningsum(nums)
	if err != nil {
		t.Fatalf("runningsum(%v) returned error: %v", nums, err)
	}
	if len(got) != len(want) {
		t.Fatalf("runningsum(%v) = %v, want %v", nums, got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("runningsum(%v) = %v, want %v", nums, got, want)
		}
	}
}

func TestRunningSumExample2(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	want := []int{1, 2, 3, 4, 5}
	got, err := runningsum(nums)
	if err != nil {
		t.Fatalf("runningsum(%v) returned error: %v", nums, err)
	}
	if len(got) != len(want) {
		t.Fatalf("runningsum(%v) = %v, want %v", nums, got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("runningsum(%v) = %v, want %v", nums, got, want)
		}
	}
}

func TestRunningSumExample3(t *testing.T) {
	nums := []int{3, 1, 2, 10, 1}
	want := []int{3, 4, 6, 16, 17}
	got, err := runningsum(nums)
	if err != nil {
		t.Fatalf("runningsum(%v) returned error: %v", nums, err)
	}
	if len(got) != len(want) {
		t.Fatalf("runningsum(%v) = %v, want %v", nums, got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("runningsum(%v) = %v, want %v", nums, got, want)
		}
	}
}
