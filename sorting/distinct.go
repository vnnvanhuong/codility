package sorting

// Brute Force
// Create an array to contain unique elements
// return size of the array
func Distinct(A []int) int {
	result := []int{}

	for i := range A {
		found := false
		for j := range result {
			if result[j] == A[i] {
				found = true
				break
			}
		}

		if !found {
			result = append(result, A[i])
		}
	}

	return len(result)
}