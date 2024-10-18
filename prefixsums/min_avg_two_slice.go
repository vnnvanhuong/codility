package prefixsums

import "math"

// A slice of size greater 1 always is combine of slice size of 2 or size of 3
// Therefore we don't need to care for slice with size greater than 3

func MinAvgTwoSlice(A []int) int {
	n := len(A)
    
    var a1, a2, a3 int
    var sum2, sum3 int

    minSum2, minSum3 := math.MaxInt64, math.MaxInt64
    minIdx2, minIdx3 := -1, -1

    for i := 0; i < n -1 ; i++ {
        a1 = A[i]
        a2 = A[i+1]
        sum2 = a1 + a2
        
        if sum2 < minSum2 {
            minSum2 = sum2
            minIdx2 = i
        }
        
        if i < n -2 {
            a3 = A[i+2]
            sum3 = a1 + a2 + a3
            if sum3 < minSum3 {
                minSum3 = sum3
                minIdx3 = i
            }
        }
    }

    if minIdx3 == -1 {
        return minIdx2
    }

    var avg2, avg3 float64
    avg2 = float64(minSum2)/float64(2)
    avg3 = float64(minSum3)/float64(3)

    if avg2 < avg3 {
        return minIdx2
    }

    if avg2 > avg3 {
        return minIdx3
    }

    if minIdx2 < minIdx3 {
        return minIdx2
    }

    return minIdx3
}