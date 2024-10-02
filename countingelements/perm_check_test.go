package countingelements

import "testing"


func Test_PermCheck(t *testing.T) {
	if r := PermCheck([]int{4, 1, 3, 2}); r != 1 {
		t.Errorf("Expected: 1, Got: %d", r)
	}

	if r := PermCheck([]int{4, 1, 3}); r != 0 {
		t.Errorf("Expected: 0, Got: %d", r)
	}
}