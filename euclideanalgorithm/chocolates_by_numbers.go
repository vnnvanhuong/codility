package euclideanalgorithm

// 10, 4
// 0 1 2 3 4 5 6 7 8 9
// 0, 4, 8, 2, 6
//
// brute force
// hashmap
// put the element to the hashmap if condition meets
// x := 0
// x = (x + M) % N
// if map contains x, stop
func ChocolatesByNumbers(N, M int) int {
	return N / greatestCommonDivisor(N, M)
}

func greatestCommonDivisor(N, M int) int {
	if N%M == 0 {
		return M
	}

	return greatestCommonDivisor(M, N%M)
}
