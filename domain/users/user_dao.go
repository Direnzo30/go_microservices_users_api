package users

import (
	"fmt"

	"github.com/Direnzo30/go_microservices_users_api/datasources/psql"
	"github.com/Direnzo30/go_microservices_users_api/utils/errors"
)

const (
	insertQuery   = "INSERT INTO users (first_name, last_name, email, username, created_at) VALUES ($1, $2, lower($3), $4, NOW()) RETURNING id, created_at;"
	retrieveQuery = "SELECT first_name, last_name, email, username, created_at FROM users WHERE id = $1"
	updateQuery   = "UPDATE users SET first_name = $1, last_name = $2, email = lower($3), username = $4 WHERE id = $5"
)

// Get performs a find by id for users
func (u *User) Get() *errors.RestError {
	var err error
	// Prepare insertion statement
	statement, err := psql.Connections.Users.Prepare(retrieveQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	// Make sure to close statement
	defer statement.Close()
	// Assign values
	err = statement.QueryRow(u.ID).Scan(&u.FirstName, &u.LastName, &u.Email, &u.Username, &u.CreatedAt)
	if err != nil {
		return errors.NotFoundError(fmt.Sprintf("User with id: %d does not exist", u.ID))
	}
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

// Update ensures to update the user in the database
func (u *User) Update() *errors.RestError {
	var err error
	// Prepare insertion statement
	statement, err := psql.Connections.Users.Prepare(updateQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	// Make sure to close statement
	defer statement.Close()
	// Assign values
	_, err = statement.Exec(u.FirstName, u.LastName, u.Email, u.Username, u.ID)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}
