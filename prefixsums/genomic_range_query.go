package prefixsums

import "math"

// O(M*N)
// convert gen string to array
// CAGCCTA -> 2132241
// loop through query
// loop through range and find min in range

func GenomicRangeQuery(S string, P []int, Q []int) []int {
	result := make([]int, len(P))
	
	for i := range P {
		start := P[i]
		end := Q[i]

		min := math.MaxInt32
		for j := start; j <= end; j++ {
			r := 4
			switch S[j] {
			case 'A':
				r = 1 
			case 'C':
				r = 2
			case 'G':
				r = 3
			}

			if r < min {
				min = r
			}
		}

		result[i]= min
	}

	return result
}


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