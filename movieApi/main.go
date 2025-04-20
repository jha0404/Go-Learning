package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	movieName := "Guardians of the Galaxy Vol. 3"
	movieName2 := "Godzilla vs. Kong"
	movieName3 := "The Flash"
	var listMovie []string
	listMovie = append(listMovie, movieName, movieName2, movieName3)
	for _, movie := range listMovie {
		MovieDetailGathering(movie)
	}
}

// Movie Json data struct
type Movie struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Languages  string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"Type"`
	DVD        string `json:"DVD"`
	BoxOffice  string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website    string `json:"Website"`
	Response   string `json:"Response"`
	Rating     []any  `json:"Ratings"`
}

func MovieDetailGathering(movieName string) {
	baseURL := "http://www.omdbapi.com/"
	imdbID := "tt3896198"
	apiKey := "901f69a2"

	URL := fmt.Sprintf("%s?t=%s&i=%s&apikey=%s", baseURL, movieName, imdbID, apiKey)
	res, err := http.Get(URL)
	if err != nil {
		panic(err)

	}
	bt, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	var movie Movie
	DecodeJson(bt, &movie)
	var movies []Movie
	movies = append(movies, movie)
	r := EncodeJson(&movies)
	fmt.Println(string(r))
}

func DecodeJson(a []byte, b *Movie) {
	err := json.Unmarshal(a, b)
	if err != nil {
		panic(err)
	}
	//fmt.Println(b)
}

func EncodeJson(input *[]Movie) []byte {
	bt, err := json.MarshalIndent(input, "", "\t")
	if err != nil {
		panic(err)
	}
	return bt
}
