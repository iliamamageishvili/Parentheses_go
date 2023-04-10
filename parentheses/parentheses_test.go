package parentheses

import (
	"net/http"
	"net/http/httptest"
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

func TestGenerateSequence(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/generate?n=10", nil)

	rec := httptest.NewRecorder()

	GenerateSequenceHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, rec.Code)
	}

	expected := 10
	actual := len(rec.Body.String())

	if actual != expected {
		t.Errorf("Expected a sequence of length %d, but got a sequence of length %d", expected, actual)
	}
}
