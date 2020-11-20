package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is for services like AWS for making ping
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
