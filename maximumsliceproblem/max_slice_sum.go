package maximumsliceproblem

// Brute Force
// 3 2 -6 4 0
// i = 0, n-2
// j = i + 1, n-1
// calculate sums from i to j
// remember the max sum
func MaxSliceSum(A []int) int {
	n := len(A)

	maxSum := A[0]
	for i := 0; i < n; i++ {
		sum := A[i]
		if sum > maxSum {
			maxSum = sum
		}

		if i == n-1 {
			break
		}

		for j := i + 1; j < n; j++ {
			sum += A[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum

}
