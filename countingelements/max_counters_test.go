package countingelements

import (
	"slices"
	"testing"
)

func Test_MaxCounters(t *testing.T) {
	expected := []int{3, 2, 2, 4, 2}
	if actual := MaxCounters(5, []int{3,4,4,6,1,4,4}); !slices.Equal(actual, expected) {
		t.Errorf("Expected: %v, Got: %v", expected, actual)
	}
}