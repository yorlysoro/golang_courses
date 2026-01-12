package ransomnote

import "testing"

func TestCanConstructExamples(t *testing.T) {
	tests := []struct {
		name       string
		ransomNote string
		magazine   string
		want       bool
	}{
		{"Example1", "a", "b", false},
		{"Example2", "aa", "ab", false},
		{"Example3", "aa", "aab", true},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := canConstruct(tc.ransomNote, tc.magazine); got != tc.want {
				t.Fatalf("canConstruct(%q, %q) = %v; want %v", tc.ransomNote, tc.magazine, got, tc.want)
			}
		})
	}
}
