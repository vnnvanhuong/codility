package primenumber_test

import (
	"testing"
	"vnnvanhuong/codility/primenumber"
)

func TestPeaks(t *testing.T) {
	tcs := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "test case 1",
			input:    []int{1, 2, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2},
			expected: 3,
		},
		{
			name:     "test case 2",
			input:    []int{1, 3, 2, 1},
			expected: 1,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := primenumber.Peaks(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, actual)
			}
		})
	}
}
