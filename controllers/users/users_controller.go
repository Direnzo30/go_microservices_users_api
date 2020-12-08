package users

import (
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
		panic(err)
	}
	// Create user and save it to database
	if err = services.UsersService.CreateUser(&user); err != nil {
		panic(err)
	}
	// Render result
	if err = renderer.RenderSingleResponse(c, &user, nil, nil, http.StatusCreated); err != nil {
		panic(err)
	}
}

// Show - show
func Show(c *gin.Context) {
	var user *users.User
	var err *errors.RestError
	// Initialize user
	user, err = initalizeUser(c)
	if err != nil {
		panic(err)
	}
	// Render result
	if err = renderer.RenderSingleResponse(c, user, nil, nil, http.StatusOK); err != nil {
		panic(err)
	}
}

// Update - update
func Update(c *gin.Context) {
	var user *users.User
	var err *errors.RestError
	// Initialize user
	user, err = initalizeUser(c)
	if err != nil {
		panic(err)
	}
	// Obtain fields to update
	if err = decoder.DecodeBodyFromJSON(c.Request, user); err != nil {
		panic(err)
	}
	// Try to update
	if err = services.UsersService.UpdateUser(user); err != nil {
		panic(err)
	}
	// Render result
	if err = renderer.RenderSingleResponse(c, user, nil, nil, http.StatusOK); err != nil {
		panic(err)
	}

}

// Delete - delete
func Delete(c *gin.Context) {
	var user *users.User
	var err *errors.RestError
	// Initialize user
	user, err = initalizeUser(c)
	if err != nil {
		panic(err)
	}
	// Try to retrieve user
	if err = services.UsersService.DeleteUser(user); err != nil {
		panic(err)
	}
	// Render result
	response := map[string]bool{"deleted": true}
	if err = renderer.RenderSingleResponse(c, response, nil, nil, http.StatusOK); err != nil {
		panic(err)
	}
}

// Search - search
func Search(c *gin.Context) {

}

// Private Methods

// initializes user for manipulation
func initalizeUser(c *gin.Context) (*users.User, *errors.RestError) {
	// Try to get id
	userID, parseError := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseError != nil {
		return nil, errors.BadRequestError("id must be a number")
	}
	// Try to retrieve user
	user, err := services.UsersService.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
