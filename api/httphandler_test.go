package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

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
