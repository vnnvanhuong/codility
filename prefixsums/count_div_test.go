package prefixsums

import "testing"

func TestCountDiv(t *testing.T) {
	if r := CountDiv(6,11,2); r != 3 {
		t.Errorf("Expected: 3, Got: %d", r)
	}

	if r := CountDiv(6,11,3); r != 2 {
		t.Errorf("Expected: 2, Got: %d", r)
	}
}