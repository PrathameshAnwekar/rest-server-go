package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)

		log.Printf("Request: %s %s from %s completed in %d ns",
			r.Method, r.URL.Path, r.RemoteAddr, time.Since(startTime).Nanoseconds())
	})
}
