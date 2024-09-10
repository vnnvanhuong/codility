package array

import (
	"testing"
)

func TestOddOccurrencesInArray(t *testing.T) {

	if a := OddOccurrencesInArray([]int{9, 3, 9, 3, 9, 7, 9}); a != 7 {
		t.Errorf("Expected: 7, Got: %d", a)
	}
}
