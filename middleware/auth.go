package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, ok := c.Get("app")
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Context not set up correctly"})
			return
		}

		app, ok := ctx.(*App)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to assert app as *App type"})
			return
		}

		redisClient := app.Redis

		token := c.Request.Header.Get("Authorization")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, no token found"})
			return
		}

		_, err := redisClient.Client.Get(token).Result()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized, expired token"})
				return
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		c.Next()
	}
}
