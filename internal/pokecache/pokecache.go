package pokecache

import (
	"time"
	"sync"
	"fmt"
	"encoding/json"
)

type cacheEntry struct {
	createdAt	time.Time //time entry was created
	val			[]byte //raw data we're caching
}

type Cache struct {
	cachedEntries	map[string]cacheEntry
	interval		time.Duration
	mu				*sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	var addEntry cacheEntry
	if err := json.Unmarshal(val, &addEntry); err != nil {
		fmt.Println(err)
	}

	c.cachedEntries[key] = addEntry

}

func (c *Cache) Get(key string) (entryData []byte, found bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.cachedEntries[key]
	if !ok {
		return nil, false
	}

	data, err := json.Marshal(val)
	if err != nil {
		return nil, false
	}
	return data, true
}

func (c *Cache) reapLoop() {
	
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		select {
		case <- ticker.C:
			c.reap()
		}
	}
}

func (c *Cache) reap() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, val := range c.cachedEntries {
		if time.Since(val.createdAt) > c.interval {
			delete(c.cachedEntries, key)
		}
	}
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
        cachedEntries: make(map[string]cacheEntry),
        interval:      interval,
        mu:			   &sync.Mutex{},
    }

	go cache.reapLoop()

	return cache

}