package main

import (
	"Parentheses_go/api"
	"Parentheses_go/parentheses"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	router := http.NewServeMux()

	router.HandleFunc("/generate", api.GenerateSequenceHandler)

	go func() {
		if err := http.ListenAndServe(":8080", router); err != nil {
			panic(err)
		}
	}()

	lengths := []int{2, 4, 8}

	for _, length := range lengths {
		balancedCount := 0

		for i := 0; i < 1000; i++ {
			resp, err := http.Get(fmt.Sprintf("http://localhost:8080/generate?n=%d", length))
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

		percentBalanced := float64(balancedCount) / 10
		fmt.Printf("Length %d: %.2f%% balanced\n", length, percentBalanced)
	}

	time.Sleep(time.Second)
}
