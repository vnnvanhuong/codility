package array

func OddOccurrencesInArray(a []int) int {
	n := len(a)
	r := 0

	for i := 0; i < n; i++ {
		r ^= a[i]
	}

	return r
}
