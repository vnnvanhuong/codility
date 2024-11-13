package primenumber_test

import (
	"testing"
	"vnnvanhuong/codility/primenumber"
)

func TestFlags(t *testing.T) {
	tcs := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "test case 1",
			input:    []int{1, 5, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2},
			expected: 3,
		},
		{
			name:     "test case 2",
			input:    []int{1},
			expected: 0,
		},
		{
			name:     "test case 3",
			input:    []int{1, 5},
			expected: 0,
		},
		{
			name:     "test case 4",
			input:    []int{1, 5, 3},
			expected: 1,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := primenumber.Flags(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, actual)
			}
		})
	}
}
