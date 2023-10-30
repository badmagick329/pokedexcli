package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) []byte {
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
	v, ok := c.cache[key]
	if ok {
		return v.val
	}
	return nil
}

func (c *Cache) Get(key string) []byte {
	v, ok := c.cache[key]
	if ok {
		return v.val
	}
	return nil
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	before := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(before) {
			delete(c.cache, k)
		}
	}
}
