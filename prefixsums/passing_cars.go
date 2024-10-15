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

func PassingCars(A []int) int {
	count := 0
	for i := 0; i < len(A) - 1 ; i++ {
		if A[i] == 0 {
			for j := i + 1; j < len(A); j++ {
				if A[j] == 1 {
					count++
				}
			}
		} 
	}

	if count > 1000000000 {
		return -1
	}

	return count
}