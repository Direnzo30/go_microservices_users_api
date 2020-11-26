package userdb

import (
	"database/sql"
	"fmt"
	"os"

	// required for psql connection
	_ "github.com/lib/pq"
)

// InitUsersDB Connect to users database
func InitUsersDB() *sql.DB {
	var client *sql.DB
	var err error
	user := os.Getenv("DB_USER")
	pssw := os.Getenv("DB_PSSW")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	dataSoruceName := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, pssw, host, port, name)

	client, err = sql.Open("postgres", dataSoruceName)
	if err != nil {
		panic(err)
	}
	if err = client.Ping(); err != nil {
		panic(err)
	}
	return client
}
