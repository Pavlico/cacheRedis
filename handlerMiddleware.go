package cacheRedis

import (
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

func HandlerCacheMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("id")
		data, err := GetByKey(key)
		rww := NewResponseWriterWrapper(w)
		switch {
		case err == redis.Nil:
			handler.ServeHTTP(rww, r)
			defer func() {
				if *rww.statusCode == http.StatusOK {
					SaveToCache(key, rww.GetBody())
				}
			}()
			return
		case err != nil:
			log.Println(err)
			handler.ServeHTTP(rww, r)
			return
		}
		rww.Write([]byte(data))
	}
}
