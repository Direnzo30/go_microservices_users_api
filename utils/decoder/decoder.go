package decoder

import (
	"encoding/json"
	"net/http"

	"github.com/Direnzo30/go_microservices_users_api/utils/errors"
)

// DecodeBodyFromJSON extracts json from request Body
// Note: with gin, c.ShouldBindJSON(&object) can be used to simulate this
func DecodeBodyFromJSON(r *http.Request, o interface{}) *errors.RestError {
	if r.Header.Get("Content-Type") != "application/json" {
		return errors.BadRequestError("Request Content-Type must be application/json")
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(o); err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}
