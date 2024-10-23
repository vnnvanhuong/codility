package sorting

import (
	"fmt"
)

// Approach 1: Brute Force
func NumberOfDiscIntersections(A []int) int {
	n := len(A)
	slices := make([][]int, n)

	for i := range n {
		slices[i] = make([]int, 2)
	}

	for i := range A {
		slices[i][0] = i - A[i]
		slices[i][1] = i + A[i]
	}

	counter := make(map[string]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}

			isIntersect := false
			for p := slices[i][0]; p <= slices[i][1]; p++ {
				for q := slices[j][0]; q <= slices[j][1]; q++ {
					if p == q {
						isIntersect = true
						pair := fmt.Sprintf("%d_%d", i, j)
						revertPair := fmt.Sprintf("%d_%d", j, i)
						if counter[pair] == 0 &&
							counter[revertPair] == 0 {
							counter[pair]++
						}
						break
					}
				}

				if isIntersect {
					break
				}
			}

		}
	}

	result := len(counter)

	if result > 10000000 {
		return -1
	}

	return result
}
