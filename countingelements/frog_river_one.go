package countingelements

// create a map contains step from 1 to X
// loop through A and remove same element from S
// if S is empty, return index of element of A
func FrogRiverOne(X int, A []int) int {
	M := make(map[int]bool)

	for i := 1; i <= X ; i++ {
		M[i]=true
	}

	for i, a := range A {
		if M[a] {
			delete(M, a)
		}

		if len(M) == 0 {
			return i
		}
	}

	return -1
}
