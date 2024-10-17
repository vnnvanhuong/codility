package prefixsums

// O(M*N)
// convert gen string to array
// CAGCCTA -> 2132241
// loop through query
// loop through range and find min in range

// CAGCCTA
// 2,5,0
// 4,5,6

// 2,4 GCC -> 2
// 5,5 T -> 4
// 0,6 CAGCCTA -> 1

// A C G T - 1 2 3 4

// CAGCCTATCA
// 2,5,0
// 8,6,9

// 2,8 GCCTATC -> 1
// 5,6 TA -> 1
// 0,9 CAGCCTATCA -> 1

// O (N + M)
// Pre-compute the prefix sums of gen occurences
//  CAGCCTA
// A [0 0 1 1 1 1 1 2 2]
// C [0 1 1 1 2 3 3 3 3]
// G [0 0 0 1 1 1 1 1 1]
// loop through the query
// check from A - G (asc), return if found A[start] != A[end]
// start = P[i], end = Q[i]
// 2,5 - 2,6 - 1 1, 1 3 -> 2
// 5,5 - 5,6 - 1 1, 3 3, 11 -> 4
// 0,9 - 0,10 - 1 3 -> 1

// AA
// [0 1 2]
// 1,1 -> 1, 2 - 1 2 -> 1

func GenomicRangeQuery(S string, P []int, Q []int) []int {
	n := len(S)

	counters := make([][]int, 3)
	counters[0] = make([]int, n+1)
	counters[1] = make([]int, n+1)
	counters[2] = make([]int, n+1)

	a := 0
	c := 0
	g := 0

	for i := range S {
		switch S[i] {
		case 'A':
			a++
		case 'C':
			c++
		case 'G':
			g++
		}

		counters[0][i+1] = a
		counters[1][i+1] = c
		counters[2][i+1] = g
	}

	result := make([]int, len(P))
	for i := range P {
		start := P[i]
		end := Q[i] + 1

		r := 4
		for j := range 2 {
			if counters[j][start] != counters[j][end] {
				r = j + 1
				break
			}
		}

		result[i] = r
	}

	return result
}
