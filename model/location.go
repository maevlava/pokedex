package model

type Location struct {
	Areas []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"areas"`
	GameIndices []struct {
		GameIndex  int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Region struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"region"`
}
