package timecomplexity

func PermMissingElem(A []int) int {
	N := len(A)
	B := make([]int, N+2)

	B[0] = 0
	for i := 0; i < N; i++ {
		B[A[i]] = 1
	}

	for i := 1; i < N+2; i++ {
		if B[i] == 0 {
			return i
		}
	}

	return 0
}
