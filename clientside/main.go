package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.cgr.in/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	fmt.Println("Status Code:", statusCode) // Print the status code

	// Print the response body as a string

}
