package userdb

import (
	"database/sql"
	"os"

	// required for psql connection
	_ "github.com/lib/pq"
)

// InitUsersDB Connect to users database
func InitUsersDB() *sql.DB {
	var client *sql.DB
	var err error
	dataSoruceName := os.Getenv("DB_SOURCE_USERS")

	client, err = sql.Open("postgres", dataSoruceName)
	if err != nil {
		panic(err)
	}
	if err = client.Ping(); err != nil {
		panic(err)
	}
	return client
}
