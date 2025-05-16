// In previous lessons, we've converted response into slices of bytes, and then strings.
// Now, we decode the response data directly into a slice of structs and return that instead!
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

const url = "https://api.shodan.io/shodan/host/8.8.8.8?key=1WPtumObPQ8fUZydjqkJnOTlo6fKsX3E"

func getIssue(url string) (*Response, error) {
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

func prettyPrint(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}

func main() {
	resp, err := getIssue(url)
	if err != nil {
		log.Fatalf("error getting issue data: %v", err)
	}

	fmt.Println("----------- ways to print out the results -----------")

	// 	Use reflection to iterate over struct fields
	v := reflect.ValueOf(*resp)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}

	fmt.Println("=== Full Response Structure ===")
	prettyPrint(resp)

	// Let's also print just one data item to compare
	fmt.Println("\n=== First Data Item Only ===")
	if len(resp.Data) > 0 {
		prettyPrint(resp.Data[0])
	}
}
