package maximumsliceproblem_test

import (
	"testing"
	"vnnvanhuong/codility/maximumsliceproblem"
)

func TestMaxDoubleSliceSum(t *testing.T) {
	tcs := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "test case 1",
			input:    []int{3, 2, 6, -1, 4, 5, -1, 2},
			expected: 17,
		},
		{
			name:     "test case 2",
			input:    []int{0, 10, -1, -5, 0},
			expected: 10,
		},
		{
			name:     "test case 3",
			input:    []int{6, 1, 5, 6, 4, 2, 9, 4},
			expected: 26,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := maximumsliceproblem.MaxDoubleSliceSum(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, actual)
			}
		})
	}
}
