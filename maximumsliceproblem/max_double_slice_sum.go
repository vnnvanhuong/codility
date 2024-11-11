package maximumsliceproblem

// Prefix Sums
// 3, 2, 6, -1, 4, 5, -1, 2
// leftSums
// rightSums
// brute force to get maxSum: leftSum + rightSum
func MaxDoubleSliceSum(A []int) int {
	leftSums := make([]int, len(A))
	rightSums := make([]int, len(A))

	for i := 1; i < len(A)-1; i++ {
		leftSums[i] = max(0, max(A[i], leftSums[i-1]+A[i]))
	}

	for i := len(A) - 2; i > 0; i-- {
		rightSums[i] = max(0, max(A[i], rightSums[i+1]+A[i]))
	}

	maxSum := 0
	for i := 1; i <= len(A)-2; i++ {
		maxSum = max(maxSum, leftSums[i-1]+rightSums[i+1])
	}

	return maxSum
}
