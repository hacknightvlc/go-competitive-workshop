package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PokemonApiClient struct {
	Client *http.Client
}

func NewPokemonApiClient() *PokemonApiClient {
	p := new(PokemonApiClient)
	p.Client = &http.Client{}

	return p
}

func (r *PokemonApiClient) FindPokemon(name string) string {
	resp, err := r.Client.Get("https://pokeapi.co/api/v2/pokemon/" + name)

	if err != nil {
		return "we're doomed"
	}

	defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "failed to read a body"
	}

	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)

	if err != nil {
		return "failed to decode json"
	}

	fmt.Println(pokemon)

	return "hehe"
}
