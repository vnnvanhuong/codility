package sorting_test

import (
	"math"
	"testing"

	"vnnvanhuong/codility/sorting"
)

func TestNumberOfDiscIntersections(t *testing.T) {
	if r := sorting.NumberOfDiscIntersections([]int{1, 5, 2, 1, 4, 0}); r != 11 {
		t.Errorf("Expected: 11, Got: %d", r)
	}

	if r := sorting.NumberOfDiscIntersections([]int{}); r != 0 {
		t.Errorf("Expected: 0, Got: %d", r)
	}

	if r := sorting.NumberOfDiscIntersections([]int{math.MaxInt32, math.MaxInt32}); r != 1 {
		t.Errorf("Expected: 1, Got: %d", r)
	}
}
