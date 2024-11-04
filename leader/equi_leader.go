package leader

// HashMap, Math
// Find the leader and total occurences (T)
// start counting
// Check if left slice has a leader: left occurences / (i + 1)/2
// Check if right slice has a leader: T - left occurences / (len - i - 1)/2
// if both slices has leaders, increase the counter
func EquiLeader(A []int) int {
	n := len(A)
	counters := make(map[int]int, n)

	totalCount := 0
	candidate := -1
	for _, a := range A {
		if counters[a] == 0 {
			counters[a] = 1
			continue
		}

		counters[a]++

		if counters[a] > n/2 {
			totalCount = counters[a]
			candidate = a
		}
	}

	if candidate == -1 {
		return 0
	}

	result := 0
	leftCount := 0

	for i, a := range A {
		if a == candidate {
			leftCount++
		}

		hasLeftLeader := false
		hasRightLeader := false

		if leftCount > (i+1)/2 {
			hasLeftLeader = true
		}

		if totalCount-leftCount > (n-i-1)/2 {
			hasRightLeader = true
		}

		if hasLeftLeader && hasRightLeader {
			result++
		}
	}

	return result
}
