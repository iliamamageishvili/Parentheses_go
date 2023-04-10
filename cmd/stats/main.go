package main

import (
	"Parentheses_go/parentheses"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	length := flag.Int("length", 2, "Length of generated sequences")
	iterations := flag.Int("iterations", 1000, "Number of iterations to run")

	serverHost := flag.String("server", "http://localhost:8080", "URL of the server to use")
	flag.Parse()

	if *length <= 0 {
		fmt.Println("Invalid length, must be positive")
		return
	}

	if *iterations <= 0 {
		fmt.Println("Invalid number of iterations, must be positive")
		return
	}

	balancedCount := 0
	for i := 0; i < *iterations; i++ {
		resp, err := http.Get(fmt.Sprintf("%s/generate?n=%d", *serverHost, *length))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		defer resp.Body.Close()

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		balanced := parentheses.IsBalanced(string(bodyBytes))
		if balanced {
			balancedCount++
		}
	}

	percentBalanced := float64(balancedCount) / float64(*iterations) * 100
	fmt.Printf("Length %d: %.2f%% balanced\n", *length, percentBalanced)
}
