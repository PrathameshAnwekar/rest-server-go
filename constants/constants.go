package constants

import "time"

const (
	DefaultPort         = 8080
	DefaultReadTimeout  = 10 * time.Second
	DefaultWriteTimeout = 10 * time.Second
	DefaultIdleTimeout  = 30 * time.Second
)

const (
	DBHost     = "rest-server-go-postgresql"
	DBPort     = 5432
	DBUsername = "postgres"
	DBPassword = ""
	DBName     = "rest_server_go"
)

const (
	RedisHost = "rest-server-go-redis"
	RedisPort = 6379
)

const (
	RectBegin      = 50
	RectEnd        = 200
	FrameThickness = 2
)
