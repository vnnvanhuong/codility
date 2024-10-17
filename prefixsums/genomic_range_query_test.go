package prefixsums

import (
	"reflect"
	"testing"
)

func TestGenomicRangeQuery(t *testing.T) {
	if r := GenomicRangeQuery("CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6}); !reflect.DeepEqual(r, []int{2, 4, 1}) {
		t.Errorf("Expected: [2,4,1], Got: %v", r)
	}

	if r := GenomicRangeQuery("CAGCCTATCA", []int{2, 5, 0}, []int{8, 6, 9}); !reflect.DeepEqual(r, []int{1, 1, 1}) {
		t.Errorf("Expected: [1,1,1], Got: %v", r)
	}

	if r := GenomicRangeQuery("AAA", []int{1, 1, 0}, []int{1, 2, 1}); !reflect.DeepEqual(r, []int{1, 1, 1}) {
		t.Errorf("Expected: [1,1,1], Got: %v", r)
	}
}
