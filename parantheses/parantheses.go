package parantheses

func isBalanced(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{'(': ')', '[': ']', '{': '}'}

	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else if c == ')' || c == ']' || c == '}' {
			if len(stack) == 0 {
				return false
			}

			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if mapping[last] != c {
				return false
			}
		}
	}

	return len(stack) == 0
}
