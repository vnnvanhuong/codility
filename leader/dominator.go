package leader

// HashMap
// Keep track the occurences of an element
// Return if element occurences > len(A)/2
func Dominator(A []int) int {
	n := len(A)
	if n == 1 {
		return 0
	}

	counters := make(map[int]int, n)

	for i, a := range A {
		if counters[a] == 0 {
			counters[a] = 1
			continue
		}

		counters[a]++

		if counters[a] > n/2 {
			return i
		}
	}

	return -1
}
