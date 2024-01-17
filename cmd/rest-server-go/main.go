package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/PrathameshAnwekar/rest-server-go/api"
	"github.com/PrathameshAnwekar/rest-server-go/constants"
	"github.com/PrathameshAnwekar/rest-server-go/db"
	"github.com/PrathameshAnwekar/rest-server-go/middleware"
)

func main() {
	var port int
	flag.IntVar(&port, "port", constants.DefaultPort, "the port to run the server on")
	serverAddress := fmt.Sprintf(":%d", port)

	database := db.NewDB()
	redisClient := db.NewRedisClient()

	server := &http.Server{
		Addr:              serverAddress,
		Handler:           setupHandlers(database, redisClient),
		ReadTimeout:       constants.DefaultReadTimeout,
		WriteTimeout:      constants.DefaultWriteTimeout,
		ReadHeaderTimeout: constants.DefaultReadTimeout,
	}

	log.Printf("Server is listening on %s...\n", serverAddress)
	if err := server.ListenAndServe(); err != nil {
		database.CloseDB()
		redisClient.Close()
		log.Fatalf("Error starting server: %s\n", err)
	}
}

// setupHandlers configures different handlers for different paths.
func setupHandlers(database *db.DB, redisClient *db.Redis) http.Handler {
	mux := http.NewServeMux()

	userHandler := api.UserHandler{DB: database}
	mux.HandleFunc("/user/put", userHandler.CreateUser)
	mux.HandleFunc("/user/get", userHandler.GetUser)
	mux.HandleFunc("/user/delete", userHandler.DeleteUser)
	mux.HandleFunc("/user/update", userHandler.UpdateUser)

	loggerMux := middleware.Logger(mux)
	authMux := middleware.Auth(loggerMux, redisClient)

	return authMux
}
