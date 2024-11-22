package sieveoferatosthenes

// 3, 1, 2, 3, 6
//
// brute force
// count divisor of each, result = len(A) - count
// time complexity: O(N*N)
//
// Sieve
// 3, 1, 2, 3, 6
// occurences 3:2, 1:1, 2:1, 6:1
// divisors 3:1, 1:0, 2:1, 6:4
// visited 3:true 1:true 2:true 6:true
func CountNonDivisible(A []int) []int {
	result := make([]int, len(A))

	occurences := make(map[int]int)
	divisors := make(map[int]int)
	visited := make(map[int]bool)

	for _, x := range A {
		occurences[x]++
		divisors[x]++
	}

	for i, x := range A {
		if visited[x] {
			result[i] = len(A) - divisors[x]
			continue
		}

		for k := 1; k*k <= x; k++ {
			if x%k != 0 {
				continue
			}

			if occurences[k] > 0 && x != 1 {
				divisors[x] += occurences[k]
			}

			if occurences[x/k] > 0 && k != 1 && k != x/k {
				divisors[x] += occurences[x/k]
			}
		}

		result[i] = len(A) - divisors[x]
		visited[x] = true
	}

	return result
}
