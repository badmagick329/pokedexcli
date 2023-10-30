package pokecache

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type Cache struct {
	Cache map[string]cacheEntry `json:"cache"`
	mu    *sync.RWMutex
}

type cacheEntry struct {
	Val       []byte    `json:"val"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		Cache: make(map[string]cacheEntry),
		mu:    &sync.RWMutex{},
	}
	err := c.Load()
	if err != nil {
		fmt.Println("Error loading cache:", err)
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) []byte {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Cache[key] = cacheEntry{
		Val:       val,
		CreatedAt: time.Now().UTC(),
	}
	v, ok := c.Cache[key]
	if ok {
		return v.Val
	}
	return nil
}

func (c *Cache) Get(key string) []byte {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.Cache[key]
	if ok {
		return v.Val
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
	c.mu.Lock()
	defer c.mu.Unlock()
	before := time.Now().UTC().Add(-interval)
	for k, v := range c.Cache {
		if v.CreatedAt.Before(before) {
			delete(c.Cache, k)
		}
	}
}

func (c *Cache) Save() error {
	jsonStr, err := json.Marshal(c.Cache)
	if err != nil {
		return err
	}
	err = os.WriteFile("cachedata.json", jsonStr, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (c *Cache) Load() error {
	data, err := os.ReadFile("cachedata.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &c.Cache)
	if err != nil {
		return err
	}
	return nil
}
