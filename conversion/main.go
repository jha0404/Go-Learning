package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Example usage of the conversion package
	// This is where you would call the conversion functions
	// and handle the results as needed.
	fmt.Println("Please enter your number which you want to add:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("There are some error in reading your input")
	} else {
		fmt.Println("You entered:", input)
		fmt.Printf("Your number is of Type %T\n", input)
		input = strings.TrimSpace(input)
		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("There are some error in converting your input")
		} else {
			fmt.Printf("Your number is of Type %T\n", value)
			number := value + 1
			fmt.Print(" new number is: ", number, "\n")

		}
	}
}
