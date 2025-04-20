package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Matches struct {
	Message string `json:"msg"`
	Status  bool   `json:"status"`
	Data    []any  `json:"data"`
}

func main() {
	GetMatches()
}

func GetMatches() {

	// create a new request using http package
	url := "https://cricket-live-line1.p.rapidapi.com/liveMatches"

	req, _ := http.NewRequest("GET", url, nil)
	// add request headers
	req.Header.Add("x-rapidapi-key", "bc8bc943e9mshfcbe413b016f47ap1d2fdajsn2b06a9acb430")
	req.Header.Add("x-rapidapi-host", "cricket-live-line1.p.rapidapi.com")

	// send request
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var b Matches
	DecodeJson(body, &b)
	for _, value := range b.Data {
		var p map[string]any
		p = value.(map[string]any)
		if p["series"] == "Indian Premier League 2025" {
			var team1, team2 string
			team1 = p["team_a"].(string)
			team2 = p["team_b"].(string)
			fmt.Printf("%v vs %v\n", team1, team2)
			var score []any
			score = p["team_a_scores_over"].([]any)
			for _, v := range score {
				//fmt.Printf("%v : %v\n", k, v)
				for key1, value1 := range v.(map[string]any) {
					ram := key1
					ram1 := value1.(string)
					fmt.Printf("%v : %v\n", ram, ram1)
				}
			}
		}
	}

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
