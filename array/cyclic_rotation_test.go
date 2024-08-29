package array

import (
	"reflect"
	"testing"
)

func TestCycliRotation(t *testing.T) {

	if a := CyclicRotation([]int{3, 8, 9, 7, 6}, 3); !reflect.DeepEqual(a, []int{9, 7, 6, 3, 8}) {
		t.Errorf("Expected: [9 7 6 3 8], Got: %v", a)
	}

	if a := CyclicRotation([]int{0, 0, 0}, 3); !reflect.DeepEqual(a, []int{0, 0, 0}) {
		t.Errorf("Expected: [0 0 0], Got: %v", a)
	}

	if a := CyclicRotation([]int{1, 2, 3, 4}, 4); !reflect.DeepEqual(a, []int{1, 2, 3, 4}) {
		t.Errorf("Expected: [1 2 3 4], Got: %v", a)
	}

	if a := CyclicRotation([]int{}, 4); !reflect.DeepEqual(a, []int{}) {
		t.Errorf("Expected: [], Got: %v", a)
	}
}
