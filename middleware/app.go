package middleware

import (
	"github.com/PrathameshAnwekar/rest-server-go/db"
	"github.com/gin-gonic/gin"
)

type App struct {
	DB    *db.DB
	Redis *db.Redis
}

func Setup(server *gin.Engine, database *db.DB, redisClient *db.Redis) {
	server.Use(func(c *gin.Context) {
		app := &App{
			DB:    database,
			Redis: redisClient,
		}
		c.Set("app", app)
		c.Next()
	})

	server.Use(gin.Recovery())
	server.Use(Logger())
	// server.Use(Auth())
}
