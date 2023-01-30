package database

import (
	"database/sql"
	"strings"

	_ "github.com/lib/pq"
)

func Server() {
	connection := strings.Join(
		([]string{
			"host=localhost",
			"port=5432",
			"user=postgres",
			"password=postgres",
			"dbname=database",
			"sslmode=disable",
		}),
		" ")
	database, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	defer database.Close()
}
