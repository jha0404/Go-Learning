package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Declaring and initliazin a variable

	resp, err := http.Get("https://google.com/")
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	fmt.Printf(" Response type is %T\n", resp)
	fmt.Printf("Respone body type is %T\n", resp.Body)

}
