package fibonacci_test

import (
	"slices"
	"testing"
	"vnnvanhuong/codility/fibonacci"
)

func TestLadder(t *testing.T) {
	tcs := []struct {
		name     string
		A        []int
		B        []int
		expected []int
	}{
		{
			name:     "test case 1",
			A:        []int{4, 4, 5, 5, 1},
			B:        []int{3, 2, 4, 3, 1},
			expected: []int{5, 1, 8, 0, 1},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := fibonacci.Ladder(tc.A, tc.B)
			if !slices.Equal(actual, tc.expected) {
				t.Errorf("Expected: %d, Got: %d", tc.expected, actual)
			}
		})
	}
}
