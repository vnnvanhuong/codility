package countingelements

// Create a hashmap contains element from 1 to len(A)
// Loop through element of A
// Remove element in hashmap if found
// Return first element of hashmap

func MissingInteger(A []int) int {
	m := make(map[int]bool)

	for i := 1; i <= len(A)+1; i++ {
		m[i]=true
	}

	for i := range A {
		if m[A[i]] {
			delete(m, A[i])
		}
	}

	for i := range m {
		return i
	}

	return 1 
}