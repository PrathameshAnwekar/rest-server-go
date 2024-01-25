package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/PrathameshAnwekar/rest-server-go/api"
	"github.com/PrathameshAnwekar/rest-server-go/constants"
	"github.com/PrathameshAnwekar/rest-server-go/db"
	"github.com/PrathameshAnwekar/rest-server-go/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	var port int
	flag.IntVar(&port, "port", constants.DefaultPort, "the port to run the server on")
	serverAddress := fmt.Sprintf(":%d", port)

	server := gin.Default()
	database := db.NewDB()
	redisClient := db.NewRedisClient()

	middleware.Setup(server, database, redisClient)
	setupRoutes(server, database)

	log.Printf("Server is listening on %s...\n", serverAddress)
	if err := server.Run(serverAddress); err != nil {
		database.CloseDB()
		redisClient.Close()
		log.Fatalf("Error starting server: %s\n", err)
	}
}

func setupRoutes(server *gin.Engine, database *db.DB) *gin.Engine {
	userHandler := api.UserHandler{DB: database}
	mediaHandler := api.MediaHandler{}

	userHandler.RegisterRoutes(server)
	mediaHandler.RegisterRoutes(server)

	return server
}
