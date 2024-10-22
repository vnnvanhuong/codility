package sorting

// Approach 1: Brute Force
func Triangle(A []int) int {
	for p := range A {
		for q := range A {
			for r := range A {
				if p < q && q < r {
					if A[p]+A[q] > A[r] &&
						A[q]+A[r] > A[p] &&
						A[p]+A[r] > A[q] {
						return 1
					}
				}

			}
		}
	}
	return 0
}
