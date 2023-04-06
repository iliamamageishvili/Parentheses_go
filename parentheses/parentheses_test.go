package parentheses

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

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
		actual := IsBalanced(tc.input)

		if actual != tc.expected {
			t.Errorf("isBalanced(%q) == %v, expected %v", tc.input, actual, tc.expected)
		}
	}
}

func TestGenerateParantheses(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/generate?n=10", nil)

	rr := httptest.NewRecorder()

	GenerateParantheses(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	body := rr.Body.String()

	if len(body) != 10 {
		t.Errorf("Handler returned wrong body length: got %v want %v", len(body), 10)
	}

	if strings.ContainsAny(body, "()") != true {
		t.Errorf("Handler returned invalid characters in body: got %v", body)
	}
}
