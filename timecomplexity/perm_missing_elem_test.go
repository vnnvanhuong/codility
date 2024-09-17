package timecomplexity

import "testing"

func Test_PermMissingElem(t *testing.T) {
	if r := PermMissingElem([]int{1, 2, 3, 5}); r != 4 {
		t.Errorf("Expected: 4, Got: %d", r)
	}
}
