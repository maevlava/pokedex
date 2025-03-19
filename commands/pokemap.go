package commands

import (
	"encoding/json"
	"fmt"
	"github.com/maevlava/pokedex/internal/pokecache"
	"github.com/maevlava/pokedex/model"
	"net/http"
	"time"
)

type PokeMapCommand struct {
	PokeMaps  []model.Location
	Page      int
	TotalPage int
	Cache     *pokecache.Cache
}
type locationResponse struct {
	Next    string `json:"next"`
	Results []struct {
		URL string `json:"url"`
	} `json:"results"`
}

func (n *PokeMapCommand) Name() string {
	return "map"
}

func (n *PokeMapCommand) Description() string {
	return "Return maps in pokemon"
}

func (n *PokeMapCommand) Execute(...string) error {
	if len(n.PokeMaps) == 0 {
		return fmt.Errorf("no maps loaded")
	}

	start := n.Page * 20
	end := start + 20
	if end > len(n.PokeMaps) {
		end = len(n.PokeMaps)
	}
	currentMaps := n.PokeMaps[start:end]

	fmt.Printf("--- Page %d/%d ---\n", n.Page+1, n.TotalPage)
	for _, m := range currentMaps {
		fmt.Printf("Location: %s (ID: %d)\n", m.Name, m.ID)
	}

	n.Page = (n.Page + 1) % n.TotalPage
	return nil
}

func LoadMap(cache *pokecache.Cache) *PokeMapCommand {
	p := &PokeMapCommand{
		PokeMaps: make([]model.Location, 0),
		Page:     0,
		Cache:    cache,
	}

	client := &http.Client{}
	nextURL := "https://pokeapi.co/api/v2/location-area/"

	for len(p.PokeMaps) < 60 && nextURL != "" {
		locations, newNextURL, err := fetchLocationURLs(client, nextURL, cache)
		if err != nil {
			fmt.Println("Error fetching location URLs:", err)
			break
		}
		nextURL = newNextURL

		for _, locationURL := range locations {
			if len(p.PokeMaps) >= 60 {
				break
			}
			loc, err := fetchLocation(client, locationURL, cache)
			if err != nil {
				fmt.Println("Error fetching location:", err)
				continue
			}
			p.PokeMaps = append(p.PokeMaps, loc)
			time.Sleep(250 * time.Millisecond)
		}
	}

	p.TotalPage = (len(p.PokeMaps) + 19) / 20
	return p
}
func fetchLocationURLs(client *http.Client, url string, cache *pokecache.Cache) ([]string, string, error) {
	if cachedData, found := cache.Get(url); found {
		var result struct {
			Results []struct {
				URL string `json:"url"`
			} `json:"results"`
			Next string `json:"next"`
		}
		if err := json.Unmarshal(cachedData, &result); err != nil {
			return nil, "", err
		}
		locations := make([]string, len(result.Results))
		for i, loc := range result.Results {
			locations[i] = loc.URL
		}
		return locations, result.Next, nil
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("response error: %v", resp.StatusCode)
	}

	result := locationResponse{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, "", err
	}

	serialized, _ := json.Marshal(result)
	cache.Add(url, serialized)

	locations := make([]string, len(result.Results))
	for i, loc := range result.Results {
		locations[i] = loc.URL
	}
	return locations, result.Next, nil
}

func fetchLocation(client *http.Client, url string, cache *pokecache.Cache) (model.Location, error) {
	var loc model.Location

	if cachedData, found := cache.Get(url); found {
		fmt.Println("Cache hit for", url)
		if err := json.Unmarshal(cachedData, &loc); err != nil {
			return model.Location{}, err
		}
		return loc, nil
	}

	resp, err := client.Get(url)
	if err != nil {
		return model.Location{}, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&loc); err != nil {
		return model.Location{}, err
	}

	cachedBytes, err := json.Marshal(loc)
	if err != nil {
		return model.Location{}, err
	}
	cache.Add(url, cachedBytes)
	fmt.Println("Fetched and cached", url)

	return loc, nil
}
