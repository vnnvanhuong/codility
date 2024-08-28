package iterations

import (
	"testing"
)

func TestBinaryGap(t *testing.T) {

	if r := BinaryGap(9); r != 2 {
		t.Errorf("Expected: 2, Got: %d", r)
	}

	if r := BinaryGap(529); r != 4 {
		t.Errorf("Expected: 4, Got: %d", r)
	}

	if r := BinaryGap(20); r != 1 {
		t.Errorf("Expected: 1, Got: %d", r)
	}

	if r := BinaryGap(15); r != 0 {
		t.Errorf("Expected: 0, Got: %d", r)
	}

	if r := BinaryGap(32); r != 0 {
		t.Errorf("Expected: 0, Got: %d", r)
	}
}
