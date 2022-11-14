package redis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
)

func GetFromCache(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rdb := InitService()
		key := r.FormValue("id")
		ctx := context.Background()
		data, err := GetByKey(key, rdb.client, ctx)
		switch {
		case err == redis.Nil:
			fmt.Println("key does not exist")
		case err != nil:
			fmt.Println("Get failed", err)
		case data == "":
			fmt.Println("value is empty")
		}
		WriteResponse(&w, 200, []byte(data))
	}

}
