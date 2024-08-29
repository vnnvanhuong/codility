package array

func CyclicRotation(a []int, k int) []int {
	n := len(a)

	if n == 0 || n == k {
		return a
	}

	for i := 0; i < k; i++ {
		lastValue := a[n-1]
		for j := n - 1; j > 0; j-- {
			a[j] = a[j-1]
		}
		a[0] = lastValue
	}

	return a
}
