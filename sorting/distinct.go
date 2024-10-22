package sorting

// Brute Force O(N**2)
// Create an array to contain unique elements
// return size of the array

// HashMap O(N*log(N)) or O(N)
func Distinct(A []int) int {
	result := make(map[int]bool)

	for i := range A {
		if !result[A[i]] {
			result[A[i]]=true
		}
	}

	return len(result)
}