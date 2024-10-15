package prefixsums

import "testing"

func Test_PassingCars(t *testing.T) {
	if r := PassingCars([]int{0, 1, 0, 1, 1}); r != 5 {
		t.Errorf("Expected: 5, Got: %d", r)
	}

	if r := PassingCars([]int{0, 1, 0, 1, 1, 0, 1}); r != 8 {
		t.Errorf("Expected: 7, Got: %d", r)
	}

	if r := PassingCars([]int{0, 1, 0}); r != 1 {
		t.Errorf("Expected: 1, Got: %d", r)
	}
}