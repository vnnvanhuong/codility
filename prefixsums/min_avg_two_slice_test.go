package prefixsums_test

import (
	"testing"

	"vnnvanhuong/codility/prefixsums"
)

func TestMinAvgTwoSlice(t *testing.T) {
	if r := prefixsums.MinAvgTwoSlice([]int{4,2,2,51,5,8}); r != 1 {
		t.Errorf("Expected: 1, Got: %d", r)
	}
}