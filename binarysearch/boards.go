package binarysearch

import "fmt"

func boards(A []int, amount int) int {
	n := len(A)
	beg := 1
	end := n
	result := -1

	for beg < end {
		mid := (beg + end) / 2

		if check(A, mid) <= amount {
			end = mid - 1
			result = mid
		} else {
			beg = mid + 1
		}
	}

	return result
}

func check(A []int, size int) int {
	n := len(A)
	boards := 0
	last := -1

	for i := range n {
		if A[i] == 1 && last < i {
			boards += 1
			last = i + size - 1
		}
	}

	return boards
}

func RunBoards() {
	fmt.Println("Expected: 12 - Got: ", boards([]int{0,1,1,0,0,0,0,1,0,0,0,0,1}, 1))
	fmt.Println("Expected: 7 - Got: ", boards([]int{0,1,1,0,0,0,0,1,0,0,0,0,1}, 2))
	fmt.Println("Expected: 3 - Got: ", boards([]int{0,1,1,0,0,0,0,1,0,0,0,0,1}, 3))
}