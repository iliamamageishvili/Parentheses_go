package parantheses

import "testing"

func TestIsBalanced(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"()[]{}", true},
		{"([])", true},
		{"((()))", true},
		{"[(])", false},
		{"{[{}]()}", true},
		{"((1 + 2) * 3) - 4)/5", false},
		{"((1 + 2) * 3) - 4/5", true},
		{"(", false},
	}

	for _, tc := range testCases {
		actual := isBalanced(tc.input)

		if actual != tc.expected {
			t.Errorf("isBalanced(%q) == %v, expected %v", tc.input, actual, tc.expected)
		}
	}
}
