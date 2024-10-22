package sorting

import "sort"

// Brute Force O(N**2)
// Create an array to contain unique elements
// return size of the array

// HashMap O(N*log(N)) or O(N)

// Sorting O(N)
func Distinct(A []int) int {
	if len(A) == 0 {
		return 0
	}

	sort.Ints(A)

	current := A[0]
	count := 1

	for i := 1; i < len(A); i++ {
		if current == A[i] {
			continue
		}

		current = A[i]
		count++
	}

	return count
}
