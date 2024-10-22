package sorting_test

import (
	"testing"

	"vnnvanhuong/codility/sorting"
)

func TestMaxProductOfThree(t *testing.T) {
	if r := sorting.MaxProductOfThree([]int{-3,1,2,-2,6,5}); r != 60 {
		t.Errorf("Expected: 60, Got: %d", r)
	}

	if r := sorting.MaxProductOfThree([]int{10,10,10}); r != 1000 {
		t.Errorf("Expected: 1000, Got: %d", r)
	}
}