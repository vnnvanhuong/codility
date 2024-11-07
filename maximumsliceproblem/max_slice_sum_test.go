package maximumsliceproblem_test

import (
	"testing"
	"vnnvanhuong/codility/maximumsliceproblem"
)

func TestMaxSliceSum(t *testing.T) {
	tcs := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "test case 1",
			input:    []int{3, 2, -6, 4, 0},
			expected: 5,
		},
		{
			name:     "test case 2",
			input:    []int{3},
			expected: 3,
		},
		{
			name:     "test case 3",
			input:    []int{3, 2, -6, 4, 0, 7},
			expected: 11,
		},
		{
			name:     "test case 4",
			input:    []int{-2, 1},
			expected: 1,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := maximumsliceproblem.MaxSliceSum(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, actual)
			}
		})
	}
}
