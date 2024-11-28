package euclideanalgorithm

// 15:75
// 15:5, 15:1
func CommonPrimeDivisors(A, B []int) int {

	count := 0
	for i := 0; i < len(A); i++ {
		if hasAllPrimeDivisors(A[i], B[i]) && hasAllPrimeDivisors(B[i], A[i]) {
			count++
		}
	}

	return count
}

func hasAllPrimeDivisors(x, y int) bool {
	if y == 1 {
		return true
	}

	gcd := greatestCommonDivisor(x, y)
	if gcd == 1 {
		return false
	}

	return hasAllPrimeDivisors(x, y/gcd)
}
