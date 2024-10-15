package prefixsums

// O(N)
// loop through from A to B, perform the division and count

func CountDiv(A, B, K int) int {
	count := 0
	for i := A ; i <= B; i++ {
		if i % K == 0 {
			count++
		}
	}

	return count
}