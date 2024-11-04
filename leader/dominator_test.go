package leader_test

import (
	"slices"
	"testing"
	"vnnvanhuong/codility/leader"
)

func TestDominator(t *testing.T) {
	tcs := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "test case 1",
			input:    []int{3, 4, 3, 2, 3 - 1, 3, 3},
			expected: []int{0, 2, 4, 6, 7},
		},
		{
			name:     "test case 2",
			input:    []int{},
			expected: []int{-1},
		},
		{
			name:     "test case 3",
			input:    []int{2147483647},
			expected: []int{0},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := leader.Dominator(tc.input)
			if !slices.Contains(tc.expected, actual) {
				t.Errorf("Expected: %d, Got: %d", tc.expected, actual)
			}
		})
	}
}
