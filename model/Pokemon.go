package model

type PokemonEncounter struct {
	Pokemon Pokemon
}
type Pokemon struct {
	Name           string `json:"name"`
	URL            string `json:"url"`
	BaseExperience int    `json:"base_experience"`
}
