package sorting_test

import (
	"testing"
	"vnnvanhuong/codility/sorting"
)

func TestTriangle(t *testing.T) {
	if r := sorting.Triangle([]int{10, 2, 5, 1, 8, 20}); r != 1 {
		t.Errorf("Expected: 1, Got: %d", r)
	}

	if r := sorting.Triangle([]int{10, 2, 5, 1}); r != 1 {
		t.Errorf("Expected: 0, Got: %d", r)
	}
}
