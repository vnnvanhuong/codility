package sieveoferatosthenes

// find primes [1..N]
// find semiprimes [1..N]
// prefixSum
// answer the query
func CountSemiPrimes(N int, P []int, Q []int) []int {
	composites := make([]bool, N+1)
	composites[0] = true
	composites[1] = true
	for i := 2; i*i <= N; i++ {
		for j := i * i; j <= N; j += i {
			if !composites[i] {
				composites[j] = true
			}
		}
	}

	semiprimes := make([]bool, N+1)
	for i := 2; i*i <= N; i++ {
		if !composites[i] {
			semiprimes[i*i] = true

			for j := i * i; j <= N; j += i {
				if !composites[j/i] {
					semiprimes[j] = true
				}
			}
		}
	}

	prefixsums := make([]int, N+1)
	for i := 1; i <= N; i++ {
		if semiprimes[i] {
			prefixsums[i] = prefixsums[i-1] + 1
			continue
		}

		prefixsums[i] = prefixsums[i-1]
	}

	result := make([]int, len(P))
	for i := 0; i < len(P); i++ {
		result[i] = prefixsums[Q[i]] - prefixsums[P[i]-1]
	}

	return result
}
