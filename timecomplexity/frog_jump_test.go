package timecomplexity

import "testing"

func TestFrogJump(t *testing.T) {
	if r := FrogJump(10, 85, 30); r != 3 {
		t.Errorf("Expected: 3, Got: %d", r)
	}
}
