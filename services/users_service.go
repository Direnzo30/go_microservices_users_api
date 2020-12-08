package services

import (
	"github.com/Direnzo30/go_microservices_users_api/domain/users"
	"github.com/Direnzo30/go_microservices_users_api/utils/errors"
)

var (
	//UsersService Handles the logic for the users service
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	GetUserByID(int64) (*users.User, *errors.RestError)
	CreateUser(*users.User) *errors.RestError
	UpdateUser(*users.User) *errors.RestError
	DeleteUser(*users.User) *errors.RestError
}

// GetUserByID handles user retrieve login
func (us *usersService) GetUserByID(userID int64) (*users.User, *errors.RestError) {
	var err *errors.RestError
	var user = users.User{ID: userID}
	if err = user.Get(); err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser Handles User creation logic
func (us *usersService) CreateUser(u *users.User) *errors.RestError {
	var err *errors.RestError
	if err = u.Validate(); err != nil {
		return err
	}
	if err = u.EncryptPassword(); err != nil {
		return err
	}
	if err = u.Save(); err != nil {
		return err
	}
	return nil
}

// UpdateUser handles User update logic
func (us *usersService) UpdateUser(u *users.User) *errors.RestError {
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
func (us *usersService) DeleteUser(u *users.User) *errors.RestError {
	if err := u.Delete(); err != nil {
		return err
	}
	return nil
}
