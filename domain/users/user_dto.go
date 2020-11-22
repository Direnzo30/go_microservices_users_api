package users

import "github.com/Direnzo30/go_microservices_users_api/utils/errors"

// User struct handles the users entity
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

// Validate control if the user istance is valid
func (u *User) Validate() *errors.RestError {
	return nil
}
