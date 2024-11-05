package maximumsliceproblem

// Take profit each day and keep track the max
// The profit can be cumulative
// 23171, 21011, 21123, 21366, 21013, 21367
//
// profit=-2160, max=0, cumulative=0
// profit=112, max=112, cummlative=112
// profit=243, max=243, cummlative=112+243
func MaxProfit(A []int) int {
	maxProfit := 0
	accProfit := 0
	profit := 0

	for i := 1; i < len(A); i++ {
		profit = A[i] - A[i-1]
		if profit < 0 {
			accProfit = 0
			continue
		}

		accProfit += profit
		if accProfit > maxProfit {
			maxProfit = accProfit
		}
	}

	return maxProfit
}
