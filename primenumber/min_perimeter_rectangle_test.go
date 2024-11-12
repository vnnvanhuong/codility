package primenumber_test

import (
	"testing"
	"vnnvanhuong/codility/primenumber"
)

func TestMinPerimeterRectangle(t *testing.T) {
	tcs := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "test case 1",
			input:    30,
			expected: 22,
		},
		{
			name:     "test case 2",
			input:    1,
			expected: 4,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := primenumber.MinPerimeterRectangle(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, actual)
			}
		})
	}
}
