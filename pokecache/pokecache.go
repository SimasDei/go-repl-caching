package pockecache

import "time"

type Cache struct {
	cache map[string]CacheEntry
}

type CacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cache: make(map[string]CacheEntry),
	}
	go cache.ReapScheduler(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = CacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.cache[key]

	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c *Cache) Reap(interval time.Duration) {
	timeThreshold := time.Now().UTC().Add(-interval)
	for key, entry := range c.cache {
		if entry.createdAt.Before(timeThreshold) {
			delete(c.cache, key)
		}
	}
}

func (c *Cache) ReapScheduler(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.Reap(interval)
	}
}
