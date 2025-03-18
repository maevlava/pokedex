package commands

import (
	"encoding/json"
	"fmt"
	"github.com/maevlava/pokedex/model"
	"net/http"
	"time"
)

type PokeMapCommand struct {
	PokeMaps  []model.Location
	Page      int
	TotalPage int
}

func LoadMap() *PokeMapCommand {
	p := &PokeMapCommand{
		PokeMaps: make([]model.Location, 0),
		Page:     0,
	}

	client := &http.Client{}
	nextURL := "https://pokeapi.co/api/v2/location-area/"

fetchLoop:
	for {
		req, err := http.NewRequest("GET", nextURL, nil)
		if err != nil {
			fmt.Printf("Error requesting pokemap: %v\n", err)
			return p
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Response error: %v\n", err)
			return p
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error reading pokemap response: %v\n", resp.StatusCode)
			resp.Body.Close()
			break
		}

		var result struct {
			Results []struct {
				URL string `json:"url"`
			} `json:"results"`
			Next string `json:"next"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			fmt.Printf("Error decoding response: %v\n", err)
			resp.Body.Close()
			break
		}
		resp.Body.Close()

		for _, location := range result.Results {
			locReq, err := http.NewRequest("GET", location.URL, nil)
			if err != nil {
				fmt.Printf("Location request error: %v\n", err)
				continue
			}
			locResp, err := client.Do(locReq)
			if err != nil {
				fmt.Printf("Location response error: %v\n", err)
				continue
			}

			var pokeMap model.Location
			if err := json.NewDecoder(locResp.Body).Decode(&pokeMap); err != nil {
				fmt.Printf("Decode error: %v\n", err)
				locResp.Body.Close()
				continue
			}
			locResp.Body.Close()

			p.PokeMaps = append(p.PokeMaps, pokeMap)

			if len(p.PokeMaps) >= 60 {
				break fetchLoop
			}
		}

		if result.Next == "" {
			break
		}
		nextURL = result.Next
		time.Sleep(250 * time.Millisecond)
	}

	return p
}

func (n *PokeMapCommand) Name() string {
	return "map"
}

func (n *PokeMapCommand) Description() string {
	return "Return maps in pokemon"
}

func (n *PokeMapCommand) Execute() error {
	loadedCount := len(n.PokeMaps)

	if loadedCount < 20 {
		return fmt.Errorf("map is not ready yet")
	}

	totalPage := (loadedCount + 19) / 20
	start := n.Page * 20
	end := start + 20
	if end > loadedCount {
		end = loadedCount
	}

	currentMaps := make([]model.Location, end-start)
	copy(currentMaps, n.PokeMaps[start:end])

	for _, m := range currentMaps {
		fmt.Printf("%s\n", m.Name)
	}

	n.Page = (n.Page + 1) % totalPage

	return nil
}
