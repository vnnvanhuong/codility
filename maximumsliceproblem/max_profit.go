package maximumsliceproblem

// There is only one selling action
// Hence, we either
// - sell a stock immediately in next day
// - or, holder a stock a sell in some next days
// for each stock price
// check the profit: price[next] - price[today]
// if profit > 0 means it can be a candidate for maxProfit
// if profit < 0 means we won't sell the stock, keep holding (profit still = 0)

func MaxProfit(A []int) int {
	maxProfit := 0
	profit := 0
	for i := 0; i < len(A)-1; i++ {
		profit += A[i+1] - A[i]

		if profit < 0 {
			profit = 0
		}

		if profit > maxProfit {
			maxProfit = profit
		}

	}

	return maxProfit
}
