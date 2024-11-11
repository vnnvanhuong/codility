package primenumber_test

import (
	"testing"
	"vnnvanhuong/codility/primenumber"
)

func TestCountFactors(t *testing.T) {
	tcs := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "test case 1",
			input:    24,
			expected: 8,
		},
		{
			name:     "test case 2",
			input:    1,
			expected: 1,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := primenumber.CountFactors(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, actual)
			}
		})
	}
}
