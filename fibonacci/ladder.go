package fibonacci

import (
	"math"
)

func Ladder(A []int, B []int) []int {
	n := len(A)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		result[i] = step(A[i]) % int(math.Pow(2, float64(B[i])))
	}

	return result
}

func step(a int) int {
	if a == 1 {
		return 1
	}

	if a == 2 {
		return 2
	}

	return step(a-1) + step(a-2)
}
