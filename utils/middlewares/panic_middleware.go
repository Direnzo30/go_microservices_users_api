package middlewares

import (
	"github.com/Direnzo30/go_microservices_users_api/utils/errors"
	"github.com/Direnzo30/go_microservices_users_api/utils/renderer"
	"github.com/gin-gonic/gin"
)

// PanicMiddleware rescues the exception in order to render response
func PanicMiddleware(c *gin.Context) {
	defer func(c *gin.Context) {
		if recoveredError := recover(); recoveredError != nil {
			switch recoveredError.(type) {
			case *errors.RestError:
				renderer.RenderRestError(c, recoveredError.(*errors.RestError))
			case error:
				renderer.RenderRestError(c, errors.InternalServerError(recoveredError.(error).Error()))
			default:
				renderer.RenderRestError(c, errors.InternalServerError(recoveredError.(string)))
			}
		}
	}(c)
	c.Next()
}
