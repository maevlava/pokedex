package model

type PokemonEncounter struct {
	Pokemon Pokemon
}
type Pokemon struct {
	Name           string         `json:"name"`
	URL            string         `json:"url"`
	BaseExperience int            `json:"base_experience"`
	Height         int            `json:"height"`
	Weight         int            `json:"weight"`
	Stats          []PokemonStats `json:"stats"`
	Types          []PokemonTypes `json:"types"`
}
type PokemonStats struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"stat"`
}
type PokemonTypes struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"type"`
}
