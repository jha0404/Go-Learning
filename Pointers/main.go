package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//variable declaration
	// way1
	var a int = 10
	var b string = "hello"
	var c float64 = 13.4

	// way 4
	h := 10
	i := "hello"

	// printing variables
	fmt.Println("value of a is ", a)
	fmt.Println("value of b is ", b)
	fmt.Printf("type of a c is %T\n", c)
	fmt.Println(" value of h and i is ", h, i)

	//user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please Enter your age:")
	age, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
	} else {
		fmt.Printf(" Type of age is %T\n", age)

	}

	// conversion of string to int
	age = strings.TrimSpace(age)
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("Error converting age to int:", err)
	} else {
		fmt.Printf("Your age is %d\n", ageInt)
	}

	// lets deep dive into pointers
	var ptr *int
	ptr = &a
	fmt.Println("value of a coming from pointer is: ", *ptr)
	fmt.Println("address of a is: ", ptr, " as well as ", &a)

}
