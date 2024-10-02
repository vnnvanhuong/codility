package countingelements

// create a hashmap contains all number from 1 to len(A)
// loop through element of A
// if the hashmap does not contain an element return 0
// else remove the corresponding element in hashmap
// at the end, if the hashmap is empty return 1
// otherwise, return 0
func PermCheck(A []int) int {
	m := make(map[int]bool, len(A))

	for i := 1; i <= len(A); i++ {
		m[i]=true
	}

	for _, a := range A {
		if !m[a] {
			return 0
		}

		delete(m, a)
	}

	if len(m) == 0 {
		return 1
	}

	return 0
}