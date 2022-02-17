package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct{}

type ListPokemonResponse struct {
	Results []Pokemon `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Stats struct {
	Percentage int  `json:"base_stat"`
	Stat       Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
}

type Sprites struct {
	FrontDefault string `json:"front_default"`
}

type PokemonInfo struct {
	BaseStat []Stats `json:"stats"`
	Sprites  Sprites `json:"sprites"`
}

const baseURL = "https://pokeapi.co/api/v2/pokemon"

func NewClient() *Client {
	return &Client{}
}

func (c *Client) ListPokemons(total int) (*ListPokemonResponse, error) {
	res, err := http.Get(fmt.Sprintf("%s?limit=%d", baseURL, total))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var listPokemon ListPokemonResponse
	json.Unmarshal(body, &listPokemon)

	return &listPokemon, nil
}

func (c *Client) GetPokemonInfo(pokemonURL string) (*PokemonInfo, error) {
	res, err := http.Get(pokemonURL)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var pokemonInfo PokemonInfo
	json.Unmarshal(body, &pokemonInfo)

	return &pokemonInfo, nil
}
