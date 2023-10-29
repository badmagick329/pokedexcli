package pokeapi

type LocationArea struct {
	Count    int                   `json:"count"`
	Next     string                `json:"next"`
	Previous any                   `json:"previous"`
	Results  []LocationAreaResults `json:"results"`
}
type LocationAreaResults struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
