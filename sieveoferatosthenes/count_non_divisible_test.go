package sieveoferatosthenes_test

import (
	"slices"
	"testing"
	"vnnvanhuong/codility/sieveoferatosthenes"
)

func TestCountNonDivisible(t *testing.T) {
	tcs := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "test case 1",
			input:    []int{3, 1, 2, 3, 6},
			expected: []int{2, 4, 3, 2, 0},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := sieveoferatosthenes.CountNonDivisible(tc.input)
			if !slices.Equal(actual, tc.expected) {
				t.Errorf("Expected: %v, Got: %v", tc.expected, actual)
			}
		})
	}
}
