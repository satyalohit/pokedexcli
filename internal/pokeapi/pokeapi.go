package pokeapi

import (
	"net/http"
	"time"

	"github.com/satyalohit/pokedexcli/internal/pokecache"
	
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}
func NewClient(cacheinterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheinterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}


