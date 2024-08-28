package iterations

import (
	"strconv"
)

func BinaryGap(N int) int {
	binStr := strconv.FormatInt(int64(N), 2)

	counter := 0
	maxCount := 0
	started := false

	for _, c := range binStr {
		if c == '1' {
			if started && counter > maxCount {
				maxCount = counter
			}

			counter = 0
			started = true
		}

		if c == '0' {
			counter++
		}
	}

	return maxCount
}
