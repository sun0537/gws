package internal

import "sync"

type Map struct {
	mu   sync.RWMutex
	data map[string]interface{}
}

func NewMap() *Map {
	return &Map{
		mu:   sync.RWMutex{},
		data: make(map[string]interface{}),
	}
}

func (c *Map) Len() int {
	c.mu.RLock()
	n := len(c.data)
	c.mu.RUnlock()
	return n
}

func (c *Map) Put(k string, v interface{}) {
	c.mu.Lock()
	c.data[k] = v
	c.mu.Unlock()
}

func (c *Map) Get(k string) (interface{}, bool) {
	c.mu.RLock()
	v, exist := c.data[k]
	c.mu.RUnlock()
	return v, exist
}

func (c *Map) Delete(k string) {
	c.mu.Lock()
	delete(c.data, k)
	c.mu.Unlock()
}

func (c *Map) Foreach(fn func(k string, v interface{})) {
	c.mu.RLock()
	for k, v := range c.data {
		fn(k, v)
	}
	c.mu.RUnlock()
}