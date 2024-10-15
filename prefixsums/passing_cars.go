package prefixsums

// 0 1 0 1 1
// 0 1, 0 3, 0 4
// 2 3, 2 4

// 0 1 0 1 1 0 1
// 01, 03, 04, 06
// 23, 24, 26
// 56

// O(N*2)
// loop through element
// if i == 0, loop through remaining elements and acount all j == 1

// O(N)
// We need to remember the number of car travelling east
// We increase the count when meeting a care travelling west
// 0 1 0 1 1
// e = 1, c = 0
// e = 1, c = 1
// e = 2, c = 1
// e = 2, c = 3
// e = 2, c = 5

func PassingCars(A []int) int {
	count := 0
	eastCars := 0
	for i := 0; i < len(A) ; i++ {
		if A[i] == 0 {
			eastCars++
			continue
		}

		count += eastCars
	}

	if count > 1000000000 {
		return -1
	}

	return count
}