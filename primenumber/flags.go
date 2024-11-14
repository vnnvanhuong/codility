package primenumber

// increase flags from 1
// try to set flags on every peak found
// if the flags is redudant means it is the result
func Flags(A []int) int {
	if len(A) < 3 {
		return 0
	}

	flags := 1
	for {
		K := flags

		for i := 1; i < len(A)-1; i++ {
			if A[i] > A[i-1] && A[i] > A[i+1] {
				K--            // set the flag
				i += flags - 1 // next peak
			}

			if K == 0 {
				break // no flag to set
			}
		}

		if K > 0 { // flag is redundant
			flags--
			break
		}

		flags++
	}

	return flags
}
