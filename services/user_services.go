package services

import (
	"github.com/Direnzo30/go_microservices_users_api/domain/users"
	"github.com/Direnzo30/go_microservices_users_api/utils/errors"
)

// GetUser handles user retrieve login
func GetUser(userID int64) (*users.User, *errors.RestError) {
	var err *errors.RestError
	var user = users.User{ID: userID}
	if err = user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser Handles User creation logic
func CreateUser(u *users.User) *errors.RestError {
	var err *errors.RestError
	if err = u.Validate(); err != nil {
		return err
	}
	if err = u.Save(); err != nil {
		return err
	}
	return nil
}

// UpdateUser handles User update logic
func UpdateUser(u *users.User) *errors.RestError {
	var err *errors.RestError
	if err = u.Validate(); err != nil {
		return err
	}
	if err = u.Update(); err != nil {
		return err
	}
	return nil
}

// DeleteUser handles User update logic
func DeleteUser(u *users.User) *errors.RestError {
	if err := u.Delete(); err != nil {
		return err
	}
	return nil
}
