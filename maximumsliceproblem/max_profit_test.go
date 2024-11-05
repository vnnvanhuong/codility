package maximumsliceproblem_test

import (
	"testing"
	"vnnvanhuong/codility/maximumsliceproblem"
)

func TestMaxProfit(t *testing.T) {
	tcs := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "test case 1",
			input:    []int{23171, 21011, 21123, 21366, 21013, 21367},
			expected: 356,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := maximumsliceproblem.MaxProfit(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, actual)
			}
		})
	}
}
