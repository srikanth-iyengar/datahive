package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func ConnectDb() *sql.DB {
	connStr := os.Getenv("PG_URI")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}
