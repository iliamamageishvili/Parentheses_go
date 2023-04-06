package main

import (
	"Parantheses_go/parantheses"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	router := http.NewServeMux()

	router.HandleFunc("/generate", parantheses.GenerateParantheses)

	http.ListenAndServe(":8080", router)
}
