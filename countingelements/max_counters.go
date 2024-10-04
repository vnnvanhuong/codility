package countingelements

// Step 1: I don't have any idea to have an effective solution at the moment
// Therefore, I start with a normal solution to solve this puzzle
// - Init an array with N
// - loop through element of A and update array N basing on the rules
// Score: 77% (Correctness: 100% - Performance: 60%)

func MaxCounters(N int, A []int) []int {
	counters := make([]int, N)

	maxCounter := 0
	for i := range A {
		if A[i] == N + 1 {
			for j := range counters {
				if counters[j] > maxCounter {
					maxCounter = counters[j]
				}
			}

			for j := range counters {
				counters[j]=maxCounter
			}

			continue
		}

		counters[A[i]-1] += 1
	}

	return counters
}