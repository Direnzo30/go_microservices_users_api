package users

import (
	"regexp"
	"strings"
	"time"

	"github.com/Direnzo30/go_microservices_users_api/utils/encrypter"

	"github.com/Direnzo30/go_microservices_users_api/utils/errors"
	"github.com/Direnzo30/go_microservices_users_api/utils/validators"
)

const (
	specialRegex = `[\!\@\#\$\_\.\-]{2}`
	lowerRegex   = `[a-z]+`
	upperRegex   = `[A-Z]+`
	numberRegex  = `[0-9]+`
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
	u.Password = strings.TrimSpace(u.Password)
	// Check password
	if err := u.validatePassword(); err != nil {
		return err
	}
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

func (u *User) validatePassword() *errors.RestError {
	if len(u.Password) < 8 {
		return errors.BadRequestError("password must have at least 8 characters")
	}
	if match, _ := regexp.MatchString(specialRegex, u.Password); !match {
		return errors.BadRequestError("password must contain 2 of the following special characters: `!` `@` `#` `$` `_` `.` `-`")
	}
	if match, _ := regexp.MatchString(lowerRegex, u.Password); !match {
		return errors.BadRequestError("password must at least one lowercase letter")
	}
	if match, _ := regexp.MatchString(upperRegex, u.Password); !match {
		return errors.BadRequestError("password must at least one capital letter")
	}
	if match, _ := regexp.MatchString(numberRegex, u.Password); !match {
		return errors.BadRequestError("password must at least one number")
	}
	return nil
}
