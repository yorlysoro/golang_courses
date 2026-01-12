package middlelinkedlist

import "testing"

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expValue *int // nil indica lista vacía
	}{
		{"odd length", []int{1, 2, 3, 4, 5}, ptr(3)},
		{"even length", []int{1, 2, 3, 4}, ptr(3)}, // segundo medio = 3
		{"single node", []int{42}, ptr(42)},
		{"empty list", []int{}, nil},
		{"two nodes", []int{7, 8}, ptr(8)}, // segundo medio = 8
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			head := BuildList(tc.input)
			mid := middleNode(head)
			if tc.expValue == nil {
				if mid != nil {
					t.Fatalf("expected nil, got %v", mid.Val)
				}
			} else {
				if mid == nil {
					t.Fatalf("expected %d, got nil", *tc.expValue)
				}
				if mid.Val != *tc.expValue {
					t.Fatalf("expected %d, got %d", *tc.expValue, mid.Val)
				}
			}
		})
	}
}

func ptr(v int) *int { return &v }
