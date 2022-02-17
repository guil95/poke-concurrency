package internal

import (
	"fmt"
	"github.com/guil95/poke-concurrency/internal/client"
	"log"
	"strings"
	"sync"
)

type UseCase struct {
	Client *client.Client
}

func NewUseCase(client *client.Client) *UseCase {
	return &UseCase{
		Client: client,
	}
}

func (uc *UseCase) GetPokemons(total int, method string) ([]*Pokemon, error) {
	pokemonsResponse, err := uc.Client.ListPokemons(total)
	if err != nil {
		return nil, err
	}

	var pokemons []*Pokemon

	for _, p := range pokemonsResponse.Results {
		pokemon := Pokemon{
			Name: strings.Title(p.Name),
			Url:  p.Url,
		}

		pokemons = append(pokemons, &pokemon)
	}

	if method == "async" {
		return uc.async(pokemons), nil
	}

	return uc.syncrono(pokemons), nil
}

func (uc *UseCase) syncrono(pokemons []*Pokemon) []*Pokemon {
	for _, p := range pokemons {
		pokemonInfo, _ := uc.Client.GetPokemonInfo(p.Url)
		p.Image = pokemonInfo.Sprites.FrontDefault
		for _, stat := range pokemonInfo.BaseStat {
			p.Stat = append(p.Stat, Stat{Name: stat.Stat.Name, Percentage: stat.Percentage})
		}
		log.Println(fmt.Sprintf("Pokemon: %s processed.", p.Name))
	}
	return pokemons
}

func (uc *UseCase) async(pokemons []*Pokemon) []*Pokemon {
	var wg sync.WaitGroup
	for _, p := range pokemons {
		wg.Add(1)

		go func(pokemon *Pokemon) {
			pokemonInfo, err := uc.Client.GetPokemonInfo(pokemon.Url)
			if err != nil {
				log.Println(err)
			}
			pokemon.Image = pokemonInfo.Sprites.FrontDefault
			for _, stat := range pokemonInfo.BaseStat {
				pokemon.Stat = append(pokemon.Stat, Stat{Name: stat.Stat.Name, Percentage: stat.Percentage})
			}

			log.Println(fmt.Sprintf("Pokemon: %s processed", pokemon.Name))

			wg.Done()
		}(p)
	}

	wg.Wait()

	return pokemons
}
