package prefixsums

// O(N)
// loop through from A to B, perform the division and count

// O(1)
// the first divisable ceil(A/K)
// remaining divisable  floor(B - ceil(A/K)*K)

// 0 1 2 3 4 5 6 7 8 9 10 11
// 6
// 10
// (10 - 6)/2 = 4
// return 4 + 1

func CountDiv(A, B, K int) int {
	// first div
    if A % K != 0 {
        A = (A - A%K) + K
    }

    // last div
    if B % K != 0 {
        B = B - B%K
    }

    midDivNums := (B - A)/K

    return midDivNums + 1 // first div
}