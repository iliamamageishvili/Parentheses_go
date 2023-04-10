package parentheses

import (
	"math/rand"
	"strings"
)

func IsBalanced(s string) bool {
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

func GenerateSequence(length int) string {
	var builder strings.Builder
	for i := 0; i < length; i++ {
		switch rand.Perm(6)[0] {
		case 0:
			builder.WriteRune('(')
		case 1:
			builder.WriteRune(')')
		case 2:
			builder.WriteRune('[')
		case 3:
			builder.WriteRune(']')
		case 4:
			builder.WriteRune('{')
		case 5:
			builder.WriteRune('}')
		}
	}
	return builder.String()
}
