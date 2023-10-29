package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	next       string
	prev       string
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		next: "",
		prev: "",
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

func (c *Client) get(url string) ([]byte, error) {
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
		return nil, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading body: %v", err)
	}
	return dat, nil
}

func (c *Client) getLocationURL(back bool) string {
	endpoint := "/location-area"
	if !back && c.next != "" {
		return c.next
	} else if back && c.prev != "" {
		return c.prev
	}
	return baseUrl + endpoint
}

func (c *Client) updateCursor(next *string, prev *string) {
	if next != nil {
		c.next = *next
	}
	if prev != nil {
		c.prev = *prev
	}
}
