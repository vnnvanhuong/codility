package sorting_test

import (
	"testing"
	"vnnvanhuong/codility/sorting"
)

func TestNumberOfDiscIntersections(t *testing.T) {
	if r := sorting.NumberOfDiscIntersections([]int{1, 5, 2, 1, 4, 0}); r != 11 {
		t.Errorf("Expected: 11, Got: %d", r)
	}
}
