package pokeapi

type LocationArea struct {
	Count    int                  `json:"count"`
	Next     *string              `json:"next"`
	Previous *string              `json:"previous"`
	Results  []LocationAreaResult `json:"results"`
}
type LocationAreaResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
