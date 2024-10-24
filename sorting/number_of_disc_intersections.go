package sorting

// Approach 1: Brute Force

// Approach 2: Counting and make the array asc in order
// 1, 5, 2, 1, 4, 0
// slices   [[0 1] [0 5] [0 4] [2 4] [0 5] [5 5]]
// 0______1______2______3______4______5
// 0______1
// 0__________________________________5
// 0___________________________4
//               2_____________4
// 0__________________________________5
//                                    5
// starts   [4 0 1 0 0 1]
//
// sorting by accumulate the previous state
// startCounters [4 4 5 5 5 6]
// at position x there is y starting state so far
// at position 0 there is 4 starting state so far
// at position 1 there is 4 starting state so far
// ...
// count the intersections
// position x has at least startCounters[x] intersections - (itseft + previous starting states)
// or more if x has end with more starting states
//
// 0 - max(4, 4) - 1 = 3
// 1 - max(4, 6) - 2 = 4
// 2 - max(5, 6) - 3 = 3
// 3 - max(5, 5) - 4 = 1
// 4 - max(5, 6) - 5 = 1
// 5 - max(6, 5) - 6 = 0
func NumberOfDiscIntersections(A []int) int {
	n := len(A)
	starts := make([]int, n)

	for i := 0; i < n; i++ {
		radius := A[i]
		startPos := i - radius
		if startPos < 0 {
			startPos = 0
		}

		starts[startPos]++
	}

	total := 0
	for i := 0; i < n; i++ {
		total += starts[i]
		starts[i] = total
	}

	interSections := 0
	for i := 0; i < n; i++ {
		radius := A[i]
		endPos := i + radius
		if endPos > n-1 {
			endPos = n - 1
		}

		m := starts[i]
		if starts[i] < starts[endPos] {
			m = starts[endPos]
		}

		interSections += m - (i + 1)

		if interSections > 10000000 {
			return -1
		}
	}

	return interSections
}
