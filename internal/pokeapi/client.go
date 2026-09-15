package pokeapi

import (
	"net/http"
	"time"

	"github.com/keepo-dot/pokedex-cli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	pokecache  *pokecache.Cache
}

func NewClient(timeout time.Duration, cache *pokecache.Cache) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokecache: cache,
	}
}
