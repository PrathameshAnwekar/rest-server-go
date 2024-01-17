package middleware

import (
	"errors"
	"net/http"

	"github.com/PrathameshAnwekar/rest-server-go/db"
	"github.com/go-redis/redis"
)

func Auth(next http.Handler, redisClient *db.Redis) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			http.Error(w, "Unauthorized, not token found.", http.StatusUnauthorized)
			return
		}

		_, err := redisClient.Client.Get(token).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				http.Error(w, "Unauthorized, expired token.", http.StatusUnauthorized)
				return
			}

			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
