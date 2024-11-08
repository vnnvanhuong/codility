package maximumsliceproblem

// Brute Force
func MaxDoubleSliceSum(A []int) int {
	n := len(A)

	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + A[i]
	}

	maxSum := 0
	for left := 0; left < n-2; left++ {
		for right := left + 2; right < n; right++ {
			for middle := left + 1; middle < right; middle++ {
				leftSum := prefixSum[middle] - prefixSum[left+1]
				rightSum := prefixSum[right-1] - prefixSum[middle]
				total := leftSum + rightSum

				if total > maxSum {
					maxSum = total
				}
			}
		}
	}

	return maxSum
}
