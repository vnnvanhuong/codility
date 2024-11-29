package fibonaccinumbers_test

import (
	"testing"
	"vnnvanhuong/codility/fibonaccinumbers"
)

func TestCommonPrimeDivisors(t *testing.T) {
	tcs := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "test case 1",
			input:    []int{0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0},
			expected: 3,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := fibonaccinumbers.FibFrog(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, actual)
			}
		})
	}
}
