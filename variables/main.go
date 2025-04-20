package main

import "fmt"

func main() {
	// This is a simple Go program that prints "Hello, World!" to the console.
	fmt.Println("Hello, World!")
	// The program ends here.
	// You can add more functionality as needed.

	// working with variables
	var name string = "Abhishek Jha"
	var age int = 25
	var height float64 = 5.9
	var isStudent bool = true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Is Student:", isStudent)
	// You can also use shorthand declaration
	name1 := "John Doe"
	age1 := 30
	height1 := 6.1
	isStudent1 := false
	fmt.Println("Name:", name1)
	fmt.Println("Age:", age1)
	fmt.Println("Height:", height1)
	fmt.Println("Is Student:", isStudent1)

}
