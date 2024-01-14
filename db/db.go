package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/PrathameshAnwekar/rest-server-go/constants"
	_ "github.com/lib/pq" // for postgres DB driver
)

type DB struct {
	Conn *sql.DB
}

func NewDB() *DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		constants.DBHost, constants.DBPort, constants.DBUsername, constants.DBPassword, constants.DBName)

	postgresDB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to open database connection: %w", err))
	}
	log.Printf("Opened database connection to %s", constants.DBName)

	err = postgresDB.Ping()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to ping database: %w", err))
	}
	log.Println("Connected to postgres database:", constants.DBName)

	return &DB{Conn: postgresDB}
}

func (db *DB) CloseDB() {
	if db != nil {
		db.Conn.Close()
		log.Println("Database Connection closed.")
	}
}
