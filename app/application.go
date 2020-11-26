package app

import (
	"github.com/Direnzo30/go_microservices_users_api/datasources/psql"
	"github.com/gin-gonic/gin"
)

// This will be controlling the app routing
var (
	router = gin.Default()
)

// StartApplication runs the server
func StartApplication() {
	// Initialize connections
	psql.InitHandlers()
	// Sets Routes for controllers
	mapUrls()
	// Launch Server
	router.Run(":3000")
}
