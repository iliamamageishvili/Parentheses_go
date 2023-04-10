package main

import (
	"Parentheses_go/api"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	router := http.NewServeMux()

	router.HandleFunc("/generate", api.GenerateSequenceHandler)

	http.ListenAndServe(":8080", router)
}

//http://localhost:8080/generate?n=150 adrress
