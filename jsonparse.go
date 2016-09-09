package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// FYI - Remember to make variables start with Uppercase so they can be accessed
// outside the package
type userRecord []struct {
	// These fields use the "json: tag" to specify which field they map to
	UserID   float64 `json:"id"`
	User     string  `json:"name"`
	UserName string  `json:"username"`
	// These fields are mapped direclty by name (note the different case)
	Email string
	Phone string
	// As these fields can be nullable, we use a pointer to a string rather
	// than a string
	Website string
}

func printHeader(resp http.Response) {
	fmt.Println("HTTP Status:", resp.Status)
	fmt.Println("HTTP Protocol:", resp.Proto)
	return
}

func getContent(url string) ([]byte, error) {
	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Send the request via a client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Print the HTTP Header
	printHeader(*resp)

	// Defer the closing of the body
	defer resp.Body.Close()

	//Read the cotnent into a byte array
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// We are done, return the bytes
	return body, nil
}

func getUserRecord(id string) (*userRecord, error) {
	content, err := getContent(fmt.Sprintf("https://jsonplaceholder.typicode.com/users?id=%s", id))
	if err != nil {
		fmt.Println("URL Error:", err)
		return nil, err
	}

	// Print the full response
	fmt.Println(string(content))

	//Fill the record with the data from the JSON
	var record userRecord
	err = json.Unmarshal(content, &record)
	if err != nil {
		fmt.Println("JSON Error:", err)
		return nil, err
	}

	return &record, err
}

func main() {
	// Simple test to get user with id=1
	record, _ := getUserRecord("1")
	fmt.Printf("User Record:\n%v\n", record)
}
