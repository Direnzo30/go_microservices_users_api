package users

import (
	"github.com/Direnzo30/go_microservices_users_api/datasources/psql"
	"github.com/Direnzo30/go_microservices_users_api/utils/errors"
)

const (
	insertQuery = "INSERT INTO users (first_name, last_name, email, username, created_at) VALUES ($1, $2, lower($3), $4, NOW()) RETURNING id, created_at;"
)

// Get performs a find by id for users
func (u *User) Get() *errors.RestError {
	return nil
}

// Save ensures to save the user in the database
func (u *User) Save() *errors.RestError {
	var err error
	// Prepare insertion statement
	statement, err := psql.Connections.Users.Prepare(insertQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	// Make sure to close statement
	defer statement.Close()
	// Assign values
	err = statement.QueryRow(u.FirstName, u.LastName, u.Email, u.Username).Scan(&u.ID, &u.CreatedAt)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}
