package users

import (
	"time"

	"github.com/Direnzo30/go_microservices_users_api/utils/errors"
	"github.com/Direnzo30/go_microservices_users_api/utils/validators"
)

// User struct handles the users entity
type User struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name" validate:"required, max=60"`
	LastName  string    `json:"last_name"  validate:"required, max=60"`
	Email     string    `json:"email"      validate:"required, email"`
	Username  string    `json:"username"   validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

// Validate control if the user istance is valid
func (u *User) Validate() *errors.RestError {
	if err := validators.ValidateStruct(u); err != nil {
		return err
	}
	return nil
}
