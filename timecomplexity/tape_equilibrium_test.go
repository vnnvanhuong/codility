package timecomplexity


import "testing"


func Test_TapEquilibrium(t *testing.T) {

	if r := TapeEquilibrium([]int{3,1,2,4,3}); r != 1 {
		t.Errorf("Expected: 1, Got: %d", r)
	}
}