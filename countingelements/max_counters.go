package countingelements

// Step 1: I don't have any idea to have an effective solution at the moment
// Therefore, I start with a normal solution to solve this puzzle
// - Init an array with N
// - loop through element of A and update array N basing on the rules
// Score: 77% (Correctness: 100% - Performance: 60%)

// Step 2: Improvement
// - Remove the nested loop to find the max value by using math.Max
// - Remove the nested loop to set all element to maxValue by storing a minValue

func MaxCounters(N int, A []int) []int {
	counters := make([]int, N)

	maxCounter := 0
	operation := 0
	minValue := 0

	for i := range A {
		operation = A[i]
		if operation == N + 1 {
			minValue = maxCounter
			continue
		}

		operation -= 1
		counters[operation] = Max(counters[operation] + 1, minValue + 1)
		maxCounter = Max(maxCounter, counters[operation])
	}

	for i := range counters {
		counters[i] = Max(minValue, counters[i])
	}

	return counters
}


func Max(x, y int) int {
	if x > y {
		return x
	}

	return y
}