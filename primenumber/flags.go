package primenumber

func Flags(A []int) int {
	n := len(A)
	if n < 3 {
		return 0
	}

	maxPeaks := len(A) / 3
	maxFlags := 0

	for i := maxPeaks; i > 0; i-- {
		flags := 0
		previousPeak := 0
		for j := 1; j < n-1; j++ {
			if A[j] > A[j-1] && A[j] > A[j+1] {
				if j == 1 || j-previousPeak >= i {
					flags++
					previousPeak = j
				}
			}
		}

		maxFlags = max(maxFlags, flags)

		if maxFlags == i {
			break
		}
	}

	return maxFlags
}

// brute force to find the maximum flags
// i from 1 to len(A)/3 since maxium peaks if len(A)/3
// j from 0 to len(A)
// max = i if there i peaks
// return max
