package sorting

import "sort"

// Approach 1: Brute Force
// Approach 2: Sorting
func Triangle(A []int) int {
	sort.Ints(A)

	if len(A) < 3 {
		return 0
	}

	for i := 0; i < len(A)-2; i++ {
		if A[i]+A[i+1] > A[i+2] {
			return 1
		}
	}
	return 0
}
