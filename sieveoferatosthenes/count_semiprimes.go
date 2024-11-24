package sieveoferatosthenes

// understand the problem
// 26
// 1:26
// 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26
// 4, 6, 9, 10, 14, 15, 21, 22, 25, 26
// 4:10
// 4 5 6 7 8 9 10
// 4 6 9 10
// 16:20
// 16 17 18 19 20
// 0
// brute force
// find primes O(NLogN) with sieve of eratosthenes
// find semiprimes O(N*N) loop through primes till the product of two < N
// query O(N*M)
func CountSemiPrimes(N int, P []int, Q []int) []int {
	composite := make([]int, N+1)
	composite[0] = 1
	composite[1] = 1
	for i := 2; i*i <= N; i++ {
		if composite[i] == 0 {
			for k := i * i; k <= N; k += i {
				composite[k] = 1
			}
		}
	}

	primes := []int{}
	for k := range composite {
		if composite[k] == 0 {
			primes = append(primes, k)
		}
	}

	result := make([]int, len(P))
	for i := 0; i < len(P); i++ {
		count := 0
		x, y := 0, 0
		for x < len(primes) && primes[x]*primes[x] <= Q[i] {
			s := primes[x] * primes[y]
			if s < P[i] || s > Q[i] || y == len(primes) {
				x++
				y = x
				continue
			}

			count++
			y++
		}

		result[i] = count
	}

	return result
}
