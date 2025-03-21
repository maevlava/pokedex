package commands

import (
	"encoding/json"
	"fmt"
	"github.com/maevlava/pokedex/model"
	"math/rand"
	"net/http"
	"time"
)

type CatchCommand struct {
}

func (e *CatchCommand) Name() string {
	return "catch"
}
func (e *CatchCommand) Description() string {
	return "Catch a Pokemon"
}
func (e *CatchCommand) Execute(user *model.User, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected 1 arguments, got %d", len(args))
	}
	// Call PokeApi for base experience of the pokemon
	pokemon, err := fetchPokemon(args[0])
	if err != nil {
		return err
	}

	// Create Rate of catching
	catchRate := getCatchRate(pokemon.BaseExperience)

	// catch success or not using random
	fmt.Printf("Throwing a Pokeball at %s...", pokemon.Name)
	isSuccess := catchPokemon(catchRate)

	// if success add it to user's Pokedex
	if isSuccess {
		user.Pokedex = append(user.Pokedex, pokemon)
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func catchPokemon(catchRate float32) bool {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomValue := r.Float32() * 100

	return catchRate >= randomValue
}
func getCatchRate(baseExperience int) float32 {
	return float32(50 - (baseExperience / 10))
}
func fetchPokemon(pokemonName string) (model.Pokemon, error) {
	//PokeApi URL =  https://pokeapi.co/api/v2/pokemon/{id or name}/
	URL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemonName)
	client := &http.Client{}
	resp, err := client.Get(URL)
	if err != nil {
		return model.Pokemon{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return model.Pokemon{}, err
	}
	pokemon := model.Pokemon{}
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return model.Pokemon{}, err
	}
	return pokemon, nil
}
