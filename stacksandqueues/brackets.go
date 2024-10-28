package stacksandqueues

// Stack
// {[()()]}
// if open bracket, push to stack
// if close bracket, pop item from stack and compare
// at the end, if len(stack) is empty, it is nested
func Brackets(s string) int {
	if len(s) == 0 {
		return 1
	}

	stack := []rune{}

	var n int
	var x rune
	for _, c := range s {
		n = len(stack)
		if n == 0 || c == '{' || c == '[' || c == '(' {
			stack = append(stack, c)
			continue
		}

		x = stack[n-1]
		stack = stack[:n-1]

		if c == '}' && x != '{' {
			return 0
		}

		if c == ']' && x != '[' {
			return 0
		}

		if c == ')' && x != '(' {
			return 0
		}
	}

	if len(stack) > 0 {
		return 0
	}

	return 1
}
