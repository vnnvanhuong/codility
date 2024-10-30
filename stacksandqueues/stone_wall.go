package stacksandqueues

// Use a Stack to compare the height of
// the previous block and current block
// use a counter for differrent blocks added to the stack
// if equal, keep the previous block, not count
// if smaller, replace the previous block, count++
// if greater, keep the previous block, count++
// return the counter
//
// for example,
// 8, 8, 5, 7, 9, 8, 7, 4, 8
// 8 - s =[8], c=1
// 8 - s =[8], c=1
// 5 - s =[5], c=2
// 7 - s =[5], c=3
// 9 - s =[5], c=4
// 8 - s =[5], c=5
// 7 - s =[5], c=5
// 4 - s =[4], c=6
// 8 - s =[4], c=7
func StoneWall(H []int) int {
	stack := []int{}
	count := 0

	for _, h := range H {
		n := len(stack)
		if n == 0 {
			stack = append(stack, h)
			count++
			continue
		}

		p := stack[n-1]
		switch {
		case h > p:
			count++
		case h < p:
			count++
			stack = stack[:n-1]
		default:
			// do nothing
		}
	}

	return count
}
