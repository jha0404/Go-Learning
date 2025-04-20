package main

import "fmt"

type Man struct {
	Name       string
	Age        int
	Complexion string
}

func (m Man) Greet() {
	fmt.Println("Hello my name is : ", m.Name)
}

func (m *Man) ChangeName() {
	m.Name = "Shyam"
}

func main() {
	var ram Man
	ram.Name = "Ram"
	ram.Age = 25
	ram.Complexion = "Fair"

	ram.Greet()
	ram.ChangeName()
	ram.Greet()

	//var intilazion and declartion
	var a int
	var b int = 10

	a = 5

	c := 10

	//slice initalization and declartion
	var slice []int // initialization
	//slice declaration
	slice = append(slice,
		1,
		2,
		3)

	//slice initilazion and declartion
	var slice2 = []int(1, 2, 3)

	s := []int{} // slice initialization and declaration
	s = append(s, 1, 2, 3)

	//slice initialization and declaration
	t := s[:1]
	s = append(s[:1], s[2:]...)

	//Map initialization and declaration
	var m map[string]int
	//map declaration
	m = make(map[string]int)
	//map initialization and declaration
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	//map initialization and declaration
	m2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	//map initialization and declaration
	m3 := map[string]int{}
	m3["one"] = 1
	m3["two"] = 2
	m3["three"] = 3
	//map initialization and declaration
	m4 := map[string]int{
		"one": 1,
	}

}
