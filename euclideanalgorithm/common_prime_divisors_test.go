package euclideanalgorithm_test

import (
	"testing"
	"vnnvanhuong/codility/euclideanalgorithm"
)

func TestCommonPrimeDivisors(t *testing.T) {
	tcs := []struct {
		name     string
		A        []int
		B        []int
		expected int
	}{
		{
			name:     "test case 1",
			A:        []int{15, 10, 3},
			B:        []int{75, 30, 5},
			expected: 1,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := euclideanalgorithm.CommonPrimeDivisors(tc.A, tc.B)
			if actual != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, actual)
			}
		})
	}
}
