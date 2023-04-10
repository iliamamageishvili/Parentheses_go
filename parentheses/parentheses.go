package parentheses

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
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
	var parentheses string
	for i := 0; i < length; i++ {
		switch rand.Perm(6)[0] {
		case 0:
			parentheses += "("
		case 1:
			parentheses += ")"
		case 2:
			parentheses += "["
		case 3:
			parentheses += "]"
		case 4:
			parentheses += "{"
		case 5:
			parentheses += "}"
		}
	}
	return parentheses
}

func GenerateSequenceHandler(w http.ResponseWriter, r *http.Request) {
	lengthParam := r.URL.Query().Get("n")

	length, err := strconv.Atoi(lengthParam)
	if err != nil {
		http.Error(w, "Invalid length parameter", http.StatusBadRequest)
		return
	}

	parentheses := GenerateSequence(length)

	fmt.Fprint(w, parentheses)
}
