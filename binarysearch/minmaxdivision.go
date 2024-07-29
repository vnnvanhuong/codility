package binarysearch

import "fmt"

func minMaxDivision(K int, M int, A []int) int {
	// what is M used for?
	// how to separate the slice a into K block?

	// how many ways to separate the slice in K block?
	// find max in k block

	// n := len(A)
	// 7, 0, 0 - 15
	// 1, 4, 2 - 9
	// 3, 0, 4 - 8
	// 2, 2, 3 - 6

	// n, 0, 0
	// n-1, 1, 0
	// n-1, 0, 1
	// n-2, 1, 1
	// n-2, 2, 0
	// n-2, 0, 2
	// min(sum(n-1), sum(1), sum(0))

	// solution: https://codility-solutions.com/lessons/lesson-14-binary-search-algorithm/minmaxdivision/
	
	return 0
}

func RunMinMaxDivision() {
	fmt.Println("Expected: 6 - Got: ", minMaxDivision(3,5,[]int{2, 1, 5, 1, 2, 2, 2}))
	fmt.Println("Expected: 3 - Got: ", minMaxDivision(2,4,[]int{2, 1, 4}))
}