package app

import (
	"errors"
	"os"

	"github.com/Direnzo30/go_microservices_users_api/datasources/psql"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// This will be controlling the app routing
var (
	router = gin.Default()
)

// StartApplication runs the server
func StartApplication() {
	// Initialize custom ENV vars
	setEnvironmentVars()
	// Initialize connections
	psql.InitHandlers()
	// Sets Routes for controllers
	mapUrls()
	// Launch Server
	router.Run(":3000")
}

func setEnvironmentVars() {
	if os.Getenv("GO_ENVIRONMET") == "development" {
		if err := godotenv.Load(".env"); err != nil {
			panic(errors.New("UNABLE TO LOAD ENV VARIABLES"))
		}
	}
}
