package validators

import (
	"fmt"

	"github.com/Direnzo30/go_microservices_users_api/utils/errors"
	"github.com/go-playground/validator"
)

// ValidateStruct apply model validators and creates generic error
func ValidateStruct(model interface{}) *errors.RestError {
	structureValidator := validator.New()
	if err := structureValidator.Struct(model); err != nil {
		firstError := err.(validator.ValidationErrors)[0]
		errorMsg := fmt.Sprintf("%s failed on %s validation", firstError.StructField(), firstError.ActualTag())
		return errors.BadRequestError(errorMsg)
	}
	return nil
}
