package main

import (
	"Parentheses_go/parentheses"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// Parse command-line arguments
	length := flag.Int("length", 2, "Length of generated sequences (2, 4, or 8)")
	iterations := flag.Int("iterations", 1000, "Number of iterations to run")
	serverHost := flag.String("server", "http://localhost:8080", "URL of the server to use")
	flag.Parse()

	// Ensure that length is valid
	validLengths := []int{2, 4, 8}
	validLength := false
	for _, l := range validLengths {
		if *length == l {
			validLength = true
			break
		}
	}
	if !validLength {
		fmt.Println("Invalid length, must be one of 2, 4, or 8")
		return
	}

	// Ensure that iterations is positive
	if *iterations <= 0 {
		fmt.Println("Invalid number of iterations, must be positive")
		return
	}

	// Seed the random number generator
	rand.Seed(time.Now().Unix())

	// Perform the iterations
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
