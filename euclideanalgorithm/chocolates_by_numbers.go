package euclideanalgorithm

// 10, 4
// 0 1 2 3 4 5 6 7 8 9
// 0, 4, 8, 2, 6
//
// brute force
// hashmap
// put the element to the hashmap if condition meets
// x := 0
// x += M % N
// if map contains x, stop
func ChocolatesByNumbers(N, M int) int {
	wrappers := make(map[int]bool)

	for x := 0; !wrappers[x]; x = (x + M) % N {
		wrappers[x] = true
	}

	return len(wrappers)
}
