package cache_fibonacci

import (
	"math/big"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	cleanupInterval time.Duration
	items           map[int64]*big.Int
}

func New(interval time.Duration) *Cache {
	cache := &Cache{
		cleanupInterval: interval,
		items:           make(map[int64]*big.Int),
	}

	if interval > 0 {
		cache.StartGC()
	}

	return cache
}

func (c *Cache) Set(key int64, value *big.Int) {
	c.Lock()
	defer c.Unlock()

	c.items[key] = value
}

func (c *Cache) Get(key int64) *big.Int {
	c.RLock()
	defer c.RUnlock()

	return c.items[key]
}

func (c *Cache) Delete(key int64) {
	c.Lock()
	defer c.Unlock()

	delete(c.items, key)
}

func (c *Cache) StartGC() {
	go c.GC()
}

func (c *Cache) GC() {
	for {
		<-time.After(c.cleanupInterval)

		if c.items == nil {
			return
		}

		c.Lock()
		c.items = make(map[int64]*big.Int)
		c.Unlock()
	}
}
