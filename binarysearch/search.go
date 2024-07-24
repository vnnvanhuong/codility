package binarysearch

import (
	"fmt"
)

func binarySearch(A []int, x int) int {
	n := len(A)

	begin := 0
	end := n-1
	result := -1

	for begin <= end {
		mid := (begin + end)/2

		if A[mid] <= x {
			begin = mid+1
			result = mid
		} else {
			end = mid - 1
		}
	}

	return result
}

func RunBinarySearch() {
	fmt.Println("Expected: 4 - Got: ", binarySearch([]int{12,15,15,19,24,31,53,59,60}, 24))
}