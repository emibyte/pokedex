package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mapping      map[string]cacheEntry
	mutex        sync.Mutex
	reapInterval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		mapping:      make(map[string]cacheEntry),
		mutex:        sync.Mutex{},
		reapInterval: interval,
	}
	go cache.reapLoop()
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	// TODO: use the read only lock here at some points since its probably better
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.mapping[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) (val []byte, found bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock() // NOTE: there was a deadlock here before since there was two returns, im dumb
	if entry, ok := c.mapping[key]; ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.reapInterval)
	defer ticker.Stop()

	for range ticker.C {
		c.mutex.Lock()
		for k, v := range c.mapping {
			if time.Since(v.createdAt) > c.reapInterval {
				delete(c.mapping, k)
			}
		}
		c.mutex.Unlock()
	}
}
