package stacksandqueues_test

import (
	"testing"
	"vnnvanhuong/codility/stacksandqueues"
)

func TestNesting(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Test case 1",
			input:    "(()(())())",
			expected: 1,
		},
		{
			name:     "Test case 2",
			input:    "())",
			expected: 0,
		},
		{
			name:     "Test case 3",
			input:    "",
			expected: 1,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if actual := stacksandqueues.Nesting(tc.input); actual != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, actual)
			}
		})
	}
}
