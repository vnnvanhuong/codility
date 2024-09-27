package timecomplexity

import "math"

// a = 3,1,2,4,3
// right = sum(a)
// left = 0
// loop through a
// right -= a[i]
// left += a[i]
// gap = MAX_INT
// if abs(right - left) < gap then assign gap
// return gap

func TapeEquilibrium(a []int) int {
	right := 0

	for i := range a {
		right += a[i]
	}

	left := 0
	gap := 0
	result := math.MaxInt

	for i := 0; i < len(a)-1; i++ {
		right -= a[i]
		left += a[i]
		gap = int(math.Abs(float64(right) - float64(left)))

		if gap < result {
			result = gap
		}
	}

	return result
}