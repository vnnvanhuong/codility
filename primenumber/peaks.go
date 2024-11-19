package primenumber

// 1, 2, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2
//
// Intituiative
//
// blockLen: 3 - N/2
// peaks: 3, 5, 10
// check if a block contains a peak
//
// Pseudo
// current = blockLen, current < N/2
// peak = 0, peak < len(peaks)
// count = 0
// if current > peak
//   move to next block
//   move to next peak
// else
//   move to next block
//
// Execution
// b=2,  2>3?   -> p=3, c=0
// b=5,  5>3?   -> p=5, c=1
// b=8,  8>5?   -> p=10, c=2
// b=11, 11>10? -> p=10, c=3
// return 3
func Peaks(A []int) int {
	N := len(A)
	if N < 3 {
		return 0
	}

	peaks := []int{}

	for i := 1; i < N-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			peaks = append(peaks, i)
		}
	}

	if len(peaks) == 0 {
		return 0
	}

	maxPeaks := 0
	for b := 3; b < N/2; b++ {
		if N%b != 0 {
			continue
		}

		p := 0
		count := 0
		for c := b - 1; c < N; c += b {
			if c >= peaks[p] {
				p++
				count++
			}
		}

		if count > maxPeaks {
			maxPeaks = count
		}
	}

	if len(peaks) > 0 && maxPeaks == 0 {
		maxPeaks = 1
	}

	return maxPeaks
}
