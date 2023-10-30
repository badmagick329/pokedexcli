package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/badmagick329/pokedexcli/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	config     Config
	cache      pokecache.Cache
}

type StatusCodeError struct {
	code int
}

func (s *StatusCodeError) Error() string {
	return fmt.Sprintf("Bad status code: %d", s.code)
}

func NewClient(cacheInterval time.Duration, config Config) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		config: config,
		cache:  pokecache.NewCache(cacheInterval),
	}
}

func (c *Client) ListLocationAreas(back bool) (LocationArea, error) {
	fullUrl := c.getLocationURL(back)
	dat, err := c.get(fullUrl)
	if err != nil {
		return LocationArea{}, err
	}
	data := LocationArea{}
	err = json.Unmarshal(dat, &data)
	if err != nil {
		return data, fmt.Errorf("Error unmarshalling: %v", err)
	}
	c.updateCursor(data.Next, data.Previous)
	return data, nil
}

func (c *Client) LocationDetails(locName string) (LocationDetails, error) {
	endpoint := "/location-area"
	fullUrl := baseUrl + endpoint + "/" + locName
	dat, err := c.get(fullUrl)
	if err != nil {
		switch e := err.(type) {
		case *StatusCodeError:
			return LocationDetails{}, fmt.Errorf("%s not found. %s", locName, e.Error())
		default:
			return LocationDetails{}, nil
		}
	}
	data := LocationDetails{}
	err = json.Unmarshal(dat, &data)
	if err != nil {
		return LocationDetails{}, err
	}
	return data, nil
}

func (c *Client) CatchPokemon(pokemon string) (Pokemon, error) {
	endpoint := "/pokemon"
	fullUrl := baseUrl + endpoint + "/" + pokemon
	dat, err := c.get(fullUrl)
	if err != nil {
		switch e := err.(type) {
		case *StatusCodeError:
			return Pokemon{}, fmt.Errorf("%s not found. %s", pokemon, e.Error())
		default:
			return Pokemon{}, nil
		}
	}
	data := Pokemon{}
	err = json.Unmarshal(dat, &data)
	if err != nil {
		return Pokemon{}, err
	}
	return data, nil
}

func (c *Client) get(url string) ([]byte, error) {
	cached := c.cache.Get(url)
	if cached != nil {
		return cached, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return nil, &StatusCodeError{code: resp.StatusCode}
	}
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading body: %v", err)
	}
	c.cache.Add(url, dat)
	return dat, nil
}

func (c *Client) getLocationURL(back bool) string {
	endpoint := "/location-area"
	if !back && c.config.next != "" {
		return c.config.next
	} else if back && c.config.prev != "" {
		return c.config.prev
	}
	return baseUrl + endpoint
}

func (c *Client) updateCursor(next *string, prev *string) {
	if next != nil {
		c.config.next = *next
	}
	if prev != nil {
		c.config.prev = *prev
	}
}
