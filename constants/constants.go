package constants

import "time"

const (
	DefaultPort         = 8080
	DefaultReadTimeout  = 10 * time.Second
	DefaultWriteTimeout = 10 * time.Second
	DefaultIdleTimeout  = 30 * time.Second
)

const (
	DBHost     = "localhost"
	DBPort     = 5432
	DBUsername = "anwprath"
	DBPassword = "anwprath"
	DBName     = "rest_server_go"
)
