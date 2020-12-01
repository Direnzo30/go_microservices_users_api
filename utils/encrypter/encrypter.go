package encrypter

import (
	"github.com/Direnzo30/go_microservices_users_api/utils/errors"
	"golang.org/x/crypto/bcrypt"
)

// SecureField encrypts a given field
func SecureField(field string) (string, *errors.RestError) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(field), 8)
	if err != nil {
		return "", errors.InternalServerError("Unable to encrypt value")
	}
	return string(encrypted), nil
}
