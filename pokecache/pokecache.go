package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache() Cache {
	return Cache{
		cache: make(map[string]cacheEntry),
	}
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
