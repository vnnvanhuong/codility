package sorting

import (
	"sort"
)

// Approach 1: Brute Force O(N**3)
// Approach 2: sort, get three last element

func MaxProductOfThree(A []int) int {
	n := len(A)
	sort.Ints(A)

	result := A[n-3] * A[n-2] * A[n-1]
	p2 := A[0] * A[1] * A[n-1]
	p3 := A[0] * A[n-2] * A[n-1]
	p4 := A[0] * A[1] * A[2]

	if p2 > result {
		result = p2
	}

	if p3 > result {
		result = p3
	}

	if p4 > result {
		result = p4
	}

	return result
}
