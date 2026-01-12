package maximumwealth

import "testing"

func TestExample1(t *testing.T) {
	accounts := [][]int{{1, 2, 3}, {3, 2, 1}}
	want := 6
	if got := maximumWealth(accounts); got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}

func TestExample2(t *testing.T) {
	accounts := [][]int{{1, 5}, {7, 3}, {3, 5}}
	want := 10
	if got := maximumWealth(accounts); got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}

func TestExample3(t *testing.T) {
	accounts := [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}
	want := 17
	if got := maximumWealth(accounts); got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}

func TestConstraintMinDimensions(t *testing.T) {
	accounts := [][]int{{1}}
	want := 1
	if got := maximumWealth(accounts); got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}

func TestConstraintMaxDimensions(t *testing.T) {
	accounts := make([][]int, 50)
	for i := 0; i < 50; i++ {
		row := make([]int, 50)
		for j := 0; j < 50; j++ {
			row[j] = 100
		}
		accounts[i] = row
	}
	// every customer has wealth 50*100 = 5000
	want := 5000
	if got := maximumWealth(accounts); got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}

func TestConstraintMinValue(t *testing.T) {
	accounts := [][]int{{1, 1, 1}}
	want := 3
	if got := maximumWealth(accounts); got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}

func TestConstraintMaxValue(t *testing.T) {
	row := make([]int, 50)
	for i := range row {
		row[i] = 100
	}
	accounts := [][]int{row}
	// single customer wealth = 50*100 = 5000
	want := 5000
	if got := maximumWealth(accounts); got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}
