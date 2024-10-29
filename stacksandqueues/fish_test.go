package stacksandqueues_test

import (
	"testing"
	"vnnvanhuong/codility/stacksandqueues"
)

func TestFish(t *testing.T) {
	tcs := []struct {
		name      string
		size      []int
		direction []int
		expected  int
	}{
		{
			name:      "test case 1",
			size:      []int{4, 3, 2, 1, 5},
			direction: []int{0, 1, 0, 0, 0},
			expected:  2,
		},
		{
			name:      "test case 2",
			size:      []int{4, 3, 2, 1, 5},
			direction: []int{1, 1, 0, 1, 0},
			expected:  1,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := stacksandqueues.Fish(tc.size, tc.direction)
			if actual != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, actual)
			}
		})
	}
}
