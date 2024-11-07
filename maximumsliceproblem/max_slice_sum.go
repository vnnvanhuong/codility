package maximumsliceproblem

// Max Slice
// 3, 2, -6, 4, 0
// m=A[0], s=A[0]
// s=max(0, max(s,s+A[i]))
// m=max(m,s)
func MaxSliceSum(A []int) int {
	maxSum := A[0]
	sum := A[0]
	for i := 1; i < len(A); i++ {
		sum = max(A[i], sum+A[i])
		maxSum = max(maxSum, sum)
	}

	return maxSum

}
