package http

import (
	"os"

	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/memcache"
)

// InitHTTP initializes the HTTP client using an appropriate cache service
func InitHTTP() {
	if memcachedURL := os.Getenv("MEMCACHE_URL"); memcachedURL != "" {
		client = httpcache.NewTransport(memcache.New(memcachedURL)).Client()
	} else {
		client = httpcache.NewTransport(httpcache.NewMemoryCache()).Client()
	}
}
