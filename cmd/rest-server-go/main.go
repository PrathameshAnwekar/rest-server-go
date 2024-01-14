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

	server := &http.Server{
		Addr:              serverAddress,
		Handler:           setupHandlers(),
		ReadTimeout:       constants.DefaultReadTimeout,
		WriteTimeout:      constants.DefaultWriteTimeout,
		ReadHeaderTimeout: constants.DefaultReadTimeout,
	}

	database := db.NewDB()

	log.Printf("Server is listening on %s...\n", serverAddress)
	if err := server.ListenAndServe(); err != nil {
		database.CloseDB()
		log.Fatalf("Error starting server: %s\n", err)
	}
}

// setupHandlers configures different handlers for different paths.
func setupHandlers() http.Handler {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/hello", api.Hello)

	loggerMux := middleware.Logger(mux)

	return loggerMux
}
