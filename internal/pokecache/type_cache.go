// Package pokecache handles all the caching logic.
package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	data      []byte
}

type Cache struct {
	entry    map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}
