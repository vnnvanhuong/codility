package sorting_test

import (
	"testing"
	"vnnvanhuong/codility/sorting"
)

func TestDistinct(t *testing.T) {
	if r := sorting.Distinct([]int{2,1,1,2,3,1}); r != 3 {
		t.Errorf("Expected: 3, Got: %d", r)
	}

	if r := sorting.Distinct([]int{1}); r != 1 {
		t.Errorf("Expected: 1, Got: %d", r)
	} 
}