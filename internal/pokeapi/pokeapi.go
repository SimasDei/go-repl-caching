package pokeapi

import (
	"net/http"
	"time"

	pockecache "github.com/SimasDei/go-repl-caching/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pockecache.Cache
	httpClient http.Client
}

func NewClient(interval time.Duration) Client {
	return Client{
		cache: pockecache.NewCache(interval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
