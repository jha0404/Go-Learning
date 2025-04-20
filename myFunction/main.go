package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(a int, b int) int {
	c := a + b
	return c
}

func multiplySlices(b *[]int) int {
	c := 1
	for i := 0; i < len(*b); i++ {
		c = c * (*b)[i]
	}
	return c
}

func st(a string) []string {
	s := []string{}
	s = append(s, a)
	return s
}
func main() {
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println("please type 1 for add,2 for subtract and 3 for string operation")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		input = strings.TrimSpace(input)
		inputInt, _ := strconv.Atoi(input)
		if inputInt == 1 {
			e := add(1, 2)
			fmt.Println("The sum is: ", e)

		} else if inputInt == 2 {

		} else {

		}

	}

}
