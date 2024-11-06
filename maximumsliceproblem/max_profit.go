package maximumsliceproblem

// Brute Force
// 0 1 2 3 4 5
// p=0, m=0

func MaxProfit(A []int) int {
	var maxProfit, profit int
	for i := 0; i < len(A)-1; i++ {
		profit = max(0, profit+A[i+1]-A[i])
		maxProfit = max(maxProfit, profit)
	}
	return maxProfit
}
