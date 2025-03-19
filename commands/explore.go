package commands

import (
	"encoding/json"
	"fmt"
	"github.com/maevlava/pokedex/model"
	"net/http"
)

// TODO: Implement Cache here
type LocationAreasResponse struct {
	ID                int                      `json:"id"`
	Name              string                   `json:"name"`
	PokemonEncounters []model.PokemonEncounter `json:"pokemon_encounters"`
}
type ExploreCommand struct {
	Pokemons      []model.Pokemon
	LocationAreas LocationAreasResponse
}

func (e *ExploreCommand) Name() string {
	return "explore"
}

func (e *ExploreCommand) Description() string {
	return "Explore the pokemons"
}

func (e *ExploreCommand) Execute(args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected 1 arguments, got %d", len(args))
	}
	client := &http.Client{}
	URL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", args[0])
	areas, err := fetchPokemonAreas(URL, client)
	if err != nil {
		return err
	}
	e.LocationAreas = areas
	for _, encounter := range areas.PokemonEncounters {
		e.Pokemons = append(e.Pokemons, encounter.Pokemon)
	}
	printPokemonEncounters(e.Pokemons)

	return nil
}

func printPokemonEncounters(pokemons []model.Pokemon) {
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemons {
		fmt.Println("- ", pokemon.Name)
	}
}

func fetchPokemonAreas(url string, client *http.Client) (LocationAreasResponse, error) {
	resp, err := client.Get(url)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return LocationAreasResponse{}, fmt.Errorf("response error: %v", resp.StatusCode)
	}

	result := LocationAreasResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return LocationAreasResponse{}, err
	}

	return result, nil
}
