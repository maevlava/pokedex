package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}
type Cache struct {
	mutex   sync.Mutex
	entries map[string]CacheEntry
	ttl     time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]CacheEntry),
		ttl:     interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.entries[key] = CacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	entry, ok := c.entries[key]
	if !ok || time.Since(entry.createdAt) > c.ttl {
		return nil, false
	}

	return entry.val, true
}
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.ttl)
	defer ticker.Stop()
	for range ticker.C {
		c.mutex.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > c.ttl {
				delete(c.entries, key)
			}
		}
		c.mutex.Unlock()
	}
}
