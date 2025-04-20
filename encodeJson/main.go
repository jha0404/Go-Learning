package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	weight  int
	Friends []string
	Address map[string]string
}
type Computer struct {
	Brand string
	Price int
	Model string
	Spec  map[string]string
	Dpr   []string
}

func encod1Json(input *[]Computer) []byte {
	bt, err := json.MarshalIndent(input, "", "\t")
	if err != nil {
		return nil
	}
	return bt

}

func main() {
	jsonInput := []Computer{
		{Brand: "Dell", Price: 50000, Model: "Inspiron", Spec: map[string]string{"ram": "8GB", "processor": "i5"}, Dpr: []string{"dell.com", "dell.in"}},
		{Brand: "HP", Price: 60000, Model: "Pavilion", Spec: map[string]string{"ram": "16GB", "processor": "i7"}, Dpr: []string{"hp.com", "hp.in"}},
		{Brand: "Lenovo", Price: 55000, Model: "ThinkPad", Spec: map[string]string{"ram": "8GB", "processor": "i5"}, Dpr: []string{"lenovo.com", "lenovo.in"}},
	}
	//fmt.Printf("jsonInput: %v\n", jsonInput)
	bt := encod1Json(&jsonInput)
	fmt.Println("Json:", string(bt))

}

// json encoder
func encodJson() {
	var p1, p2 Person
	p1.Name = "Abhishek"
	p1.Age = 23
	p1.weight = 23
	p1.Friends = []string{"Amit", "Raghav", "rave"}
	p1.Address = map[string]string{"city": "delhi", "state": "up", "country": "india"}

	p2 = Person{
		Name:    "prt",
		Age:     23,
		weight:  23,
		Friends: []string{"Amit", "Raghav", "rave"},
		Address: map[string]string{"city": "delhi", "state": "up", "country": "india"},
	}
	var data []Person
	data = append(data, p1)
	data = append(data, p2)
	// json 	encode
	bt, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Json:", string(bt))

}
