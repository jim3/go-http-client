package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const url = "https://api.shodan.io/shodan/host/IPADDR?key=APIKEY"

// getIpInfo sends an HTTP GET request to the specified URL and decodes the JSON response
// into a Response struct. It returns a pointer to the Response struct and an error if
// the request fails or the response cannot be decoded.
//
// Parameters:
//   - url: The URL to send the GET request to.
//
// Returns:
//   - *Response: Pointer to the decoded Response struct.
//   - error: An error if the request or decoding fails.
func getIpInfo(url string) (*Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	// Create a nil slice of resp
	var resp Response

	// Create a new decoder
	decoder := json.NewDecoder(res.Body)

	// Decode the response body into the resp slice
	if err := decoder.Decode(&resp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &resp, nil
}

func prettyPrint(v any) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}

func main() {
	resp, err := getIpInfo(url)
	if err != nil {
		log.Fatalf("error getting data: %v", err)
	}

	fmt.Println("=== Full Response Structure ===")
	prettyPrint(resp)

	// Let's also print just one data item to compare
	fmt.Println("\n=== First Data Item Only ===")
	if len(resp.Data) > 0 {
		prettyPrint(resp.Data[0])
	}
}
