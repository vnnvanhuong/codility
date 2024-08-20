package iterations


import (
	"testing"
	"vnnvanhuong/codility/iterations"
)


func TestBinaryGap(t *testing.T) {

	if iterations.BinaryGap(9) != 2 {
		t.Errorf("Expected: 2, Got: %d", iterations.BinaryGap(9))
	}

	if iterations.BinaryGap(529) != 4 {
		t.Errorf("Expected: 4, Got: %d", iterations.BinaryGap(529))
	}

	if iterations.BinaryGap(20) != 1 {
		t.Errorf("Expected: 1, Got: %d", iterations.BinaryGap(20))
	}

	if iterations.BinaryGap(15) != 0 {
		t.Errorf("Expected: 0, Got: %d", iterations.BinaryGap(15))
	}

	if iterations.BinaryGap(32) != 0 {
		t.Errorf("Expected: 0, Got: %d", iterations.BinaryGap(32))
	}
}