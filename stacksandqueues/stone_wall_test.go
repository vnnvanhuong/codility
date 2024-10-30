package stacksandqueues_test

import (
	"testing"
	"vnnvanhuong/codility/stacksandqueues"
)

func TestStoneWall(t *testing.T) {
	tcs := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "test case 1",
			height:   []int{8, 8, 5, 7, 9, 8, 7, 4, 8},
			expected: 7,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := stacksandqueues.StoneWall(tc.height)
			if actual != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, actual)
			}
		})
	}
}
