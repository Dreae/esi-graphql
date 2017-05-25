package http

import (
	"os"

	"github.com/garyburd/redigo/redis"
	"github.com/gregjones/httpcache"
	httpredis "github.com/gregjones/httpcache/redis"
)

// InitHTTP initializes the HTTP client using an appropriate cache service
func InitHTTP() {
	if redisURL := os.Getenv("REDIS_HOST"); redisURL != "" {

		redisPass := os.Getenv("REDIS_PASS")
		var options []redis.DialOption
		if redisPass != "" {
			options = append(options, redis.DialPassword(redisPass))
		}

		conn, err := redis.Dial("tcp", redisURL, options...)
		if err != nil {
			panic(err)
		}

		client = httpcache.NewTransport(httpredis.NewWithClient(conn)).Client()
	} else {
		client = httpcache.NewTransport(httpcache.NewMemoryCache()).Client()
	}
}
