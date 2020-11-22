package users

import "github.com/Direnzo30/go_microservices_users_api/utils/errors"

// Get performs a find by id for users
func (u *User) Get(userID int64) *errors.RestError {
	return nil
}

// Save ensures to save the user in the database
func (u *User) Save() *errors.RestError {
	return nil
}
