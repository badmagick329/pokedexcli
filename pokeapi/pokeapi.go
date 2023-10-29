package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Get(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func Json[T any](bytes []byte, data *T) {
	err := json.Unmarshal(bytes, data)
	if err != nil {
		log.Fatal(err)
	}
}

