package psql

import (
	"database/sql"

	"github.com/Direnzo30/go_microservices_users_api/datasources/psql/userdb"
)

// DBhandler Wrapper for all psql connections
type DBhandler struct {
	Users *sql.DB
}

var (
	// Connections centralizes
	Connections *DBhandler
)

// InitHandlers initializes all database connections
func InitHandlers() {
	Connections = &DBhandler{}
	Connections.Users = userdb.InitUsersDB()
}
