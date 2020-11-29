package app

import (
	"github.com/Direnzo30/go_microservices_users_api/controllers/ping"
	"github.com/Direnzo30/go_microservices_users_api/controllers/users"
	"github.com/Direnzo30/go_microservices_users_api/utils/middlewares"
)

func mapUrls() {
	v1 := router.Group("/api/v1")
	v1.Use(middlewares.PanicMiddleware)
	// Ping Controller Routes
	v1.GET("/ping", ping.Ping)
	// Users Controller Routes
	v1.GET("/users/:id", users.Show)
	//v1.GET("/users/search", controllers.SearchUser)
	v1.POST("/users", users.Create)
}
