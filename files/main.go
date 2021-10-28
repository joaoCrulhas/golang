package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//Response provided by pokeapi
type Response struct {
	Count   int       `json:"count"`
	Next    string    `json:"next"`
	Results []Pokemon `json:"results"`
}

//Pokemon with name and url
type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func main() {
	jsonFile, err := os.Open("pokemons.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened pokemons.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var response Response

	json.Unmarshal(byteValue, &response)

	for i := 0; i < len(response.Results); i++ {
		pokemon := (strings.Title(response.Results[i].Name))
		fmt.Printf("Pokemon => %v || URL => %v \n", pokemon, response.Results[i].URL)
	}
}
