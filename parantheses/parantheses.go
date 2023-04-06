package parantheses

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

func GenerateParantheses(w http.ResponseWriter, r *http.Request) {
	lengthParam := r.URL.Query().Get("n")

	length, err := strconv.Atoi(lengthParam)
	if err != nil {
		http.Error(w, "Invalid length parameter", http.StatusBadRequest)
		return
	}

	parentheses := ""
	for i := 0; i < length; i++ {
		if rand.Intn(2) == 0 {
			parentheses += "("
		} else {
			parentheses += ")"
		}
	}

	fmt.Fprint(w, parentheses)
}
