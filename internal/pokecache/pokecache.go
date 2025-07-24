package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu    sync.RWMutex
	entry map[string]cacheEntry
	ttl   time.Duration
}

func NewCache(ttl time.Duration) *Cache {
	cache := Cache{entry: make(map[string]cacheEntry), ttl: ttl}
	cache.reapLoop()
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	copiedVal := make([]byte, len(val))
	copy(copiedVal, val)
	newEntry := cacheEntry{createdAt: time.Now(), val: copiedVal}
	c.entry[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	entry, exists := c.entry[key]
	if !exists {
		return nil, false
	}

	if time.Since(entry.createdAt) > c.ttl {
		return nil, false
	}

	copiedVal := make([]byte, len(entry.val))
	copy(copiedVal, entry.val)

	return copiedVal, true
}

func (c *Cache) reapLoop() {
	go func() {
		interval := c.ttl

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			c.reap()
		}
	}()
}

func (c *Cache) reap() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.entry {
		if time.Since(entry.createdAt) > c.ttl {
			delete(c.entry, key)
		}
	}
}
