package encoder

import (
	"bytes"
	"encoding/gob"

	"github.com/Direnzo30/go_microservices_users_api/utils/errors"
)

// SerializeObject sets object properties on serializer
func SerializeObject(object interface{}, serializer interface{}) *errors.RestError {
	buffer := bytes.Buffer{}
	var err error
	err = gob.NewEncoder(&buffer).Encode(object)
	if err != nil {
		return errors.InternalServerError("Unable to serialize response")
	}
	err = gob.NewDecoder(&buffer).Decode(serializer)
	if err != nil {
		return errors.InternalServerError("Unable to serialize response")
	}
	return nil
}
