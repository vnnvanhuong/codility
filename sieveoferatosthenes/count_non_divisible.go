package sieveoferatosthenes

// 3, 1, 2, 3, 6
//
// brute force
// count divisor of each, result = len(A) - count
// time complexity: O(N*N)
//
// hashmap
// O(N)
func CountNonDivisible(A []int) []int {
	r := make([]int, len(A))

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			if A[i]%A[j] != 0 {
				r[i] += 1
			}
		}
	}

	return r
}
