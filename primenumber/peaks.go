package primenumber

import "math"

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

	isPrimeLen := true
	for i := 2; i <= int(math.Sqrt(float64(N))); i++ {
		if N%i == 0 {
			isPrimeLen = false
			break
		}
	}

	if isPrimeLen {
		for i := 1; i < N-1; i++ {
			if A[i] > A[i-1] && A[i] > A[i+1] {
				return 1
			}
		}
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

	for b := int(math.Max(float64(3), float64(N/len(peaks)))); b <= N/2; b++ {
		if N%b != 0 {
			continue
		}

		p := -1
		count := 0
		for _, peak := range peaks {
			found := peak / b
			if p < found {
				p = found
				count++
				continue
			}

			if p > found {
				break
			}
		}

		if count == N/b {
			return count
		}
	}

	return 1
}
