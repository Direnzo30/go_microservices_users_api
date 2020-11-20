package app

import (
	"github.com/gin-gonic/gin"
)

// This will be controlling the app routing
var (
	router = gin.Default()
)

// StartApplication runs the server
func StartApplication() {
	mapUrls()
	router.Run(":3000")
}
