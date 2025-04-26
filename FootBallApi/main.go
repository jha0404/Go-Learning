package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Matches struct {
	Status   string         `json:"status"`
	Response map[string]any `json:"response"`
}

func DecodeJson(a []byte, b *Matches) {
	err := json.Unmarshal(a, b)
	if err != nil {
		panic(err)
	}
	//fmt.Println(b)
}
func EncodeJson(input *[]Matches) []byte {
	bt, err := json.MarshalIndent(input, "", "\t")
	if err != nil {
		panic(err)
	}
	return bt
}
func main() {
	url := "https://free-api-live-football-data.p.rapidapi.com/football-current-live"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", "bc8bc943e9mshfcbe413b016f47ap1d2fdajsn2b06a9acb430")
	req.Header.Add("x-rapidapi-host", "free-api-live-football-data.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var b Matches
	DecodeJson(body, &b)
	//var input []Matches
	//input = append(input, b)
	//bt := EncodeJson(&input)
	//fmt.Println(string(bt))
	var response map[string]any
	response = b.Response
	liveMatches := response["live"].([]any)
	//fmt.Println(liveMatches...)
	for _, value := range liveMatches {
		var p map[string]any
		p = value.(map[string]any)
		var home map[string]any
		var away map[string]any
		home = p["home"].(map[string]any)
		away = p["away"].(map[string]any)
		match1 := home["name"].(string)
		match2 := away["name"].(string)
		fmt.Printf("%v vs %v\n", match1, match2)

	}

}
