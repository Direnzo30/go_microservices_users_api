package users

import (
	"strings"
	"time"

	"github.com/Direnzo30/go_microservices_users_api/utils/encrypter"

	"github.com/Direnzo30/go_microservices_users_api/utils/errors"
	"github.com/Direnzo30/go_microservices_users_api/utils/validators"
)

// User struct handles the users entity
type User struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"first_name" validate:"required,max=60"`
	LastName  string    `json:"last_name"  validate:"required,max=60"`
	Email     string    `json:"email"      validate:"required,email"`
	Username  string    `json:"username"   validate:"required"`
	Password  string    `json:"password"   validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Validate control if the user istance is valid
func (u *User) Validate() *errors.RestError {
	if err := validators.ValidateStruct(u); err != nil {
		return err
	}
	// Sanitize Structure
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.Email = strings.TrimSpace(u.Email)
	u.Username = strings.TrimSpace(u.Username)
	// Check uniqueness
	if err := u.checkUniqueness(); err != nil {
		return err
	}
	return nil
}

// EncryptPassword makes password storage secure
func (u *User) EncryptPassword() *errors.RestError {
	pass, err := encrypter.SecureField(u.Password)
	if err != nil {
		return errors.InternalServerError("Unable to encrypt user's password")
	}
	u.Password = pass
	return nil
}
