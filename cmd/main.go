package main

import (
	"Parentheses_go/parentheses"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	router := http.NewServeMux()

	router.HandleFunc("/generate", parentheses.GenerateSequenceHandler)

	http.ListenAndServe(":8080", router)
}

//http://localhost:8080/generate?n=150 adrress
