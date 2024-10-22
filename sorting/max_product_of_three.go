package sorting

import (
	"math"
)

// Approach 1: Brute Force O(N**3)

func MaxProductOfThree(A []int) int {
	n := len(A)
	result := math.MinInt64
	var product int

	for p := 0; p < n-2; p++ {
		for q := p + 1; q < n-1; q++ {
			for r := p + 2; r < n; r++ {
				if p < q && q < r {
					product = A[p] * A[q] * A[r]

					if product > result {
						result = product
					}
				}
			}
		}
	}

	return result
}
