package sieveoferatosthenes_test

import (
	"slices"
	"testing"
	"vnnvanhuong/codility/sieveoferatosthenes"
)

func TestCountSemiPrimes(t *testing.T) {
	tcs := []struct {
		name     string
		N        int
		P        []int
		Q        []int
		expected []int
	}{
		{
			name:     "test case 1",
			N:        26,
			P:        []int{1, 4, 16},
			Q:        []int{26, 10, 20},
			expected: []int{2, 4, 3, 2, 0},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := sieveoferatosthenes.CountSemiPrimes(
				tc.N, tc.P, tc.Q)
			if !slices.Equal(actual, tc.expected) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, actual)
			}
		})
	}
}
