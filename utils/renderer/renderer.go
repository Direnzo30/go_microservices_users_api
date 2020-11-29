package renderer

import (
	"github.com/Direnzo30/go_microservices_users_api/utils/encoder"
	"github.com/Direnzo30/go_microservices_users_api/utils/errors"
	"github.com/gin-gonic/gin"
)

// RenderSingleResponse gives a basic structure for response
func RenderSingleResponse(c *gin.Context, object interface{}, metadata interface{}, serializer interface{}, status int) *errors.RestError {
	response := make(map[string]interface{})
	// Must serialize data
	if serializer != nil {
		if err := encoder.SerializeObject(object, serializer); err != nil {
			return err
		}
		response["result"] = serializer
		// Must return raw content
	} else {
		response["result"] = object
	}
	response["metadata"] = metadata
	// this renders the response
	c.JSON(status, response)
	return nil
}

// RenderRestError handles error response in case of panic
func RenderRestError(c *gin.Context, e *errors.RestError) {
	response := make(map[string]interface{})
	response["result"] = e
	c.JSON(e.Status, response)
}
