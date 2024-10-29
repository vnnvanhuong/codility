package stacksandqueues

// 4:0 3:1 2:0 1:0 5:0
// survivals, stack
// stack empty, survivals++
// if 1 push to stack
// loop through stack
// if 0, pop the stack
// if size < pop.size
// push back the pop
// else
// survial++
// return survivals + stack.size()
// 4:0 -> 1, []
// 3:1 -> 1, [3:1]
// 2:0 -> 1, [3:1]
// 1:0 -> 1, [3:1]
// 5:0 -> 2, []
func Fish(A, B []int) int {
	survivals := 0
	stack := []int{}

	for i := 0; i < len(A); i++ {
		if B[i] == 1 {
			stack = append(stack, i)
			continue
		}

		s := len(stack)
		for s > 0 {
			f := stack[s-1]
			stack = stack[:s-1]
			s--
			if A[i] < A[f] {
				stack = append(stack, f)
				break
			}
		}

		if len(stack) == 0 {
			survivals++
		}
	}

	return survivals + len(stack)
}
