package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Weather struct {
	Request    map[string]any
	Location   map[string]any
	Current    map[string]any
	Astro      map[string]any
	AirQuality map[string]any
	Misc       map[string]any
}

func WeatherDetailGathering() {
	URL := "http://api.weatherstack.com/current?access_key=97483667f129f954d66390cbdfd4a9d0&query=Noida"
	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var nd Weather
	json.Unmarshal(content, &nd)
	fmt.Println(nd)
	prettyjson, _ := json.MarshalIndent(&nd, "", "\t")
	fmt.Println(string(prettyjson))
	//fmt.Println(string(content))
}

func main() {
	WeatherDetailGathering()

}
