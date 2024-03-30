package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Define a struct to unmarshal the API response
type ApiResponse struct {
	Results []struct {
		Name struct {
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
	} `json:"results"`
}

func main() {
	start := time.Now()

	const numRequests = 300
	var wg sync.WaitGroup
	wg.Add(numRequests)

	// Channel to receive the API responses
	responses := make(chan ApiResponse, numRequests)

	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()
			resp, err := http.Get("https://randomuser.me/api/")
			if err != nil {
				fmt.Println("Error fetching data:", err)
				return
			}
			defer resp.Body.Close()

			var apiResponse ApiResponse
			if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
				fmt.Println("Error decoding JSON:", err)
				return
			}

			responses <- apiResponse
		}()
	}

	// Close the responses channel once all goroutines have finished
	go func() {
		wg.Wait()
		close(responses)
	}()

	// Process the responses
	for response := range responses {
		fmt.Println("Fetched data for:", response.Results[0].Name.First, response.Results[0].Name.Last)
	}

	fmt.Printf("All requests finished in %s\n", time.Since(start))
}


//// GPT-4 SUCKS AT CODING

// THE HELL IS THIS??