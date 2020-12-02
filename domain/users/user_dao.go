package users

import (
	"fmt"

	"github.com/Direnzo30/go_microservices_users_api/datasources/psql"
	"github.com/Direnzo30/go_microservices_users_api/utils/errors"
)

const (
	insertQuery   = "INSERT INTO users (first_name, last_name, email, username, password, created_at, updated_at) VALUES ($1, $2, lower($3), $4, $5, NOW(), NOW()) RETURNING id, created_at, updated_at;"
	retrieveQuery = "SELECT first_name, last_name, email, username, created_at FROM users WHERE id = $1"
	updateQuery   = "UPDATE users SET first_name = $1, last_name = $2, email = lower($3), username = $4, updated_at = NOW() WHERE id = $5 RETURNING updated_at"
	deleteQuery   = "DELETE FROM users WHERE id = $1"
	uniqueQuery   = "SELECT 1 FROM users WHERE id <> $1 AND (LOWER(email) = LOWER($2) OR LOWER(username) = LOWER($3)) LIMIT 1"
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
	err = statement.QueryRow(u.FirstName, u.LastName, u.Email, u.Username, u.Password).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}

// Update ensures to update the user in the database
func (u *User) Update() *errors.RestError {
	var err error
	// Prepare update statement
	statement, err := psql.Connections.Users.Prepare(updateQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	// Make sure to close statement
	defer statement.Close()
	// Assign values
	err = statement.QueryRow(u.FirstName, u.LastName, u.Email, u.Username, u.ID).Scan(&u.UpdatedAt)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}

// Delete ensures to delete the user from the database
func (u *User) Delete() *errors.RestError {
	var err error
	// Prepare delete statement
	statement, err := psql.Connections.Users.Prepare(deleteQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	// Make sure to close statement
	defer statement.Close()
	// Assign values
	_, err = statement.Exec(u.ID)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}

// Private Methods

func (u *User) checkUniqueness() *errors.RestError {
	var err error
	// Prepare delete statement
	statement, err := psql.Connections.Users.Prepare(uniqueQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	// Make sure to close statement
	defer statement.Close()
	// Assign values
	var dummy int64
	err = statement.QueryRow(u.ID, u.Email, u.Username).Scan(&dummy)
	// Scanned can not be done so the parameters are not taken
	if err != nil {
		return nil
	}
	return errors.BadRequestError(fmt.Sprintf("Username: %s or Email: %s has already been taken", u.Username, u.Email))
}
