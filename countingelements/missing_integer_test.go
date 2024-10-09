package countingelements

import "testing"

func Test_MissingInteger(t *testing.T) {
	if r := MissingInteger([]int{1, 3, 6, 4, 1, 2}); r != 5 {
		t.Errorf("Expected: 5, Got: %d", r)
	}

	if r := MissingInteger([]int{1, 2, 3}); r != 4 {
		t.Errorf("Expected: 4, Got: %d", r)
	}

	if r := MissingInteger([]int{-1, -3}); r != 1 {
		t.Errorf("Expected: 1, Got: %d", r)
	}
}