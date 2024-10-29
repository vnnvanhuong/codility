package stacksandqueues

// (()(())())
// if open bracket, push to stack
// if closed bracket, pop the stack
// return 0 if S empty
// e.g
// ((
// (
// (((
// (
// ((
func Nesting(S string) int {
	if len(S) == 0 {
		return 1
	}

	stack := []rune{}

	for _, v := range S {
		if v == '(' {
			stack = append(stack, v)
			continue
		}

		n := len(stack)
		if v == ')' && n == 0 {
			return 0
		}

		stack = stack[:n-1]
	}

	if len(stack) == 0 {
		return 1
	}

	return 0
}
