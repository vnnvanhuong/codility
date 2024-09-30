package countingelements

import "testing"

func Test_FrogRiverOne(t *testing.T) {
	if r := FrogRiverOne(5, []int{1, 3, 1, 4, 2, 3, 5, 4}); r != 6 {
		t.Errorf("Expected: 6, Got: %d", r)
	}
}
