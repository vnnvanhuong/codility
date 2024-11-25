package euclideanalgorithm_test

import (
	"testing"
	"vnnvanhuong/codility/euclideanalgorithm"
)

func TestCountSemiPrimes(t *testing.T) {
	tcs := []struct {
		name     string
		N        int
		M        int
		expected int
	}{
		{
			name:     "test case 1",
			N:        10,
			M:        4,
			expected: 5,
		},
		{
			name:     "test case 2",
			N:        1,
			M:        4,
			expected: 1,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := euclideanalgorithm.ChocolatesByNumbers(tc.N, tc.M)
			if actual != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, actual)
			}
		})
	}
}
