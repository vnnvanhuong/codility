package leader_test

import (
	"testing"
	"vnnvanhuong/codility/leader"
)

func TestEquiLeader(t *testing.T) {
	tcs := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "test case 1",
			input:    []int{4, 3, 4, 4, 4, 2},
			expected: 2,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := leader.EquiLeader(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, actual)
			}
		})
	}
}
