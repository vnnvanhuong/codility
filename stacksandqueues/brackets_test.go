package stacksandqueues_test

import (
	"testing"
	"vnnvanhuong/codility/stacksandqueues"
)

func TestBrackets(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Nested",
			input:    "{[()()]}",
			expected: 1,
		},
		{
			name:     "Not nested 1",
			input:    "([)()]",
			expected: 0,
		},
		{
			name:     "Not nested 2",
			input:    "",
			expected: 1,
		},
		{
			name:     "Not nested 3",
			input:    "[[[[[[[",
			expected: 0,
		},
		{
			name:     "Not nested ",
			input:    "]]]]]",
			expected: 0,
		},
		{
			name:     "Not nested ",
			input:    "[[[[[[[[[]]]]]]]]]]]",
			expected: 0,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actual := stacksandqueues.Brackets(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected: %d, Got: %d", tc.expected, actual)
			}
		})
	}
}
