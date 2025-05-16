package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const url = "https://api.shodan.io/shodan/host/IPADDR?key=APIKEY"

// getIssue sends an HTTP GET request to the specified URL and decodes the JSON response
// into a Response struct. It returns a pointer to the Response struct and an error if
// the request fails or the response cannot be decoded.
//
// Parameters:
//   - url: The URL to send the GET request to.
//
// Returns:
//   - *Response: Pointer to the decoded Response struct.
//   - error: An error if the request or decoding fails.
func getIssue(url string) (*Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	// Create a struct variable
	var resp Response

	// Create a new decoder
	decoder := json.NewDecoder(res.Body)

	// Decode the response body into the struct
	if err := decoder.Decode(&resp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}
	return &resp, nil
}

func main() {
	resp, err := getIssue(url)
	if err != nil {
		log.Fatalf("error getting issue data: %v", err)
	}

	// Pretty print the response
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
