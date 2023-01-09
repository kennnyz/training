package cache

import (
	"sync"
	"time"
)

type Item struct {
	value      interface{}
	expiration time.Time
}

type Cache struct {
	myMap map[string]Item
	sync  *sync.RWMutex
}

func New() *Cache {
	cache := Cache{
		make(map[string]Item),
		new(sync.RWMutex),
	}
	go CacheClean(&cache)
	return &cache
}

func (c *Cache) Set(key string, value interface{}, deadline time.Duration) {
	c.myMap[key] = Item{
		value:      value,
		expiration: time.Now().Add(deadline),
	}
}

func (c Cache) Get(key string) interface{} {
	c.sync.RLock()
	defer c.sync.RUnlock()
	cacheItem := c.myMap[key]
	return cacheItem.value
}

func (c *Cache) Delete(key string) {
	c.sync.Lock()
	delete(c.myMap, key)
	c.sync.Unlock()
}

func CacheClean(cache *Cache) {
	for {
		for key, value := range cache.myMap {
			if isExpired(value.expiration) {
				cache.Delete(key)
			}
		}
	}
}

func isExpired(expiration time.Time) bool {
	if time.Now().After(expiration) {
		return true
	}
	return false
}
