package cache

type Cache struct {
	myMap map[string]any
}

func New() *Cache {
	return &Cache{make(map[string]any)}
}

func (c *Cache) Set(key string, value interface{}) {
	c.myMap[key] = value
}

func (c Cache) Get(key string) any {
	return c.myMap[key]
}

func (c *Cache) Delete(key string) {
	delete(c.myMap, key)
}
