package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Direnzo30/go_microservices_users_api/utils/renderer"

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
	// Obtain body
	if err = decoder.DecodeBodyFromJSON(c.Request, &user); err != nil {

	}
	// Create user and save it to database
	if err = services.CreateUser(&user); err != nil {

	}
	// Render response should go here
	if err = renderer.RenderSingleResponse(c, &user, nil, nil, http.StatusCreated); err != nil {

	}
}

// Show - show
func Show(c *gin.Context) {
	var user *users.User
	var err *errors.RestError
	userID, parseError := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseError != nil {
		//err = errors.BadRequestError("id must be a number")
	}
	user, err = services.GetUser(userID)
	if err != nil {

	}
	fmt.Printf("%+v\n", user)
}

// Search - search
func Search(c *gin.Context) {

}
