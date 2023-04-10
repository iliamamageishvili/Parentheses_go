package api

import (
	"Parentheses_go/parentheses"
	"fmt"
	"net/http"
	"strconv"
)

func GenerateSequenceHandler(w http.ResponseWriter, r *http.Request) {
	lengthParam := r.URL.Query().Get("n")

	length, err := strconv.Atoi(lengthParam)
	if err != nil {
		http.Error(w, "Invalid length parameter", http.StatusBadRequest)
		return
	}

	parentheses := parentheses.GenerateSequence(length)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, parentheses)
}
