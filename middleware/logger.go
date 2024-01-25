package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		log.Printf("%s - [%v] %s %s (%v)\n",
			c.Request.Method,
			c.Writer.Status(),
			c.Request.URL.Path,
			c.Request.Proto,
			latency,
		)
	}
}
