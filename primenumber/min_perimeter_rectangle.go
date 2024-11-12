package primenumber

import "math"

// for i=1,sqrt(N)
// find minPer of i*(n/i) if n % i = 0
func MinPerimeterRectangle(N int) int {
	minPer := math.MaxInt

	for i := 1; i*i <= N; i++ {
		if N%i == 0 {
			minPer = min(minPer, 2*(i+N/i))
		}
	}

	return minPer
}
