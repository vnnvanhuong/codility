package prefixsums

import "math"

func MinAvgTwoSlice(A []int) int {
	M := math.MaxFloat64
    r := math.MaxInt64
    
    for i := 0; i < len(A)-1; i++ {
        sum := float64(A[i])
        n := float64(1)
        for j:= i + 1; j<len(A); j++ {
            sum += float64(A[j])
            n++
            a := float64(sum/n)

            if a < M {
                M = a
                r = i
            }
        }
    }

    return r
}