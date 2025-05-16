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

	//1. Use reflection to iterate over struct fields
	v := reflect.ValueOf(*resp)
	typeOfS := v.Type() // Gets type info of a value

	// NumField returns the number of fields in the struct
	// Field returns the i'th field of the struct v
	for i := range v.NumField() {
		fmt.Printf("%s: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}

	// ----------------------------------------------

	// 2. Use MarshalIndent
	b, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
