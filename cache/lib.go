package cache

import (
	"sync"
	"time"
)

type MemCache struct {
	lock   sync.RWMutex
	store  map[interface{}]item
	maxAge int64
}

type item struct {
	created int64
	value   interface{}
}

func (cache *MemCache) Set(key interface{}, value interface{}) {
	cache.lock.Lock()
	cache.store[key] = item{
		time.Now().Unix(),
		value,
	}
	cache.lock.Unlock()
}

func (cache *MemCache) Get(key interface{}) (interface{}, bool) {
	cache.lock.RLock()
	item, ok := cache.store[key]
	cache.lock.RUnlock()

	if ok {
		if item.created < time.Now().Unix()-cache.maxAge {
			cache.lock.Lock()
			delete(cache.store, key)
			cache.lock.Unlock()

			return nil, false
		}

		return item.value, true
	}

	return nil, false
}

func New(maxAge int64) MemCache {
	return MemCache{
		store:  make(map[interface{}]item),
		maxAge: maxAge,
	}
}
