package primenumber

import (
	"math"
)

func CountFactors(N int) int {
	count := 0
	sqrt := math.Sqrt(float64(N))
	for i := 1.0; i <= sqrt; i++ {
		if i == sqrt {
			count += 1
			continue
		}

		if N%int(i) == 0 {
			count += 2
		}
	}
	return count
}
