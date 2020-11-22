package users

import (
	"github.com/Direnzo30/go_microservices_users_api/domain/users"
	"github.com/Direnzo30/go_microservices_users_api/services"
	"github.com/Direnzo30/go_microservices_users_api/utils/decoder"
	"github.com/Direnzo30/go_microservices_users_api/utils/errors"
	"github.com/gin-gonic/gin"
)

// Create - create
func Create(c *gin.Context) {
	var user users.User
	var err *errors.RestError
	if err = decoder.DecodeBodyFromJSON(c.Request, &user); err != nil {

	}
	if err = services.CreateUser(&user); err != nil {

	}
}

// Show - show
func Show(c *gin.Context) {

}

// Search - search
func Search(c *gin.Context) {

}
