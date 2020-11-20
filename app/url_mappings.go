package app

import (
	"github.com/Direnzo30/go_microservices_users_api/controllers/ping"
	"github.com/Direnzo30/go_microservices_users_api/controllers/users"
)

func mapUrls() {
	v1 := router.Group("/api/v1")
	// Ping Controller Routes
	v1.GET("/ping", ping.Ping)
	// Users Controller Routes
	v1.GET("/users/:id", users.GetUser)
	//v1.GET("/users/search", controllers.SearchUser)
	v1.POST("/users", users.CreateUser)
}
