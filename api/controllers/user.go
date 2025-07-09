package controllers

import (
	"fmt"
	"net/http"

	"github.com/OgiDac/CompanyTask/domain"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUseCase domain.UserUseCase
}

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Returns a list of all users
// @Tags         users
// @Produce      json
// @Success      200 {array} domain.User
// @Failure      400 {object} map[string]string
// @Router       /public/api/users [get]
func (uc *UserController) GetAllUsers(c *gin.Context) {
	ctx := c.Request.Context()

	users, err := uc.UserUseCase.GetAllUsers(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Registers a new user and returns access and refresh tokens
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body domain.SignUpRequest true "Sign Up Request"
// @Success      200 {object} domain.SignupResponse
// @Failure      400 {object} map[string]string
// @Router       /public/api/users [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var req domain.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "error parsing the request")
		return
	}

	ctx := c.Request.Context()
	accessToken, refreshToken, err := uc.UserUseCase.CreateUser(ctx, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, response)
}

// UpdateUser godoc
// @Summary      Update a user
// @Description  Updates user name and email
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body domain.UpdateRequest true "Update Request"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Router       /private/api/users [put]
// @Security BearerAuth
func (uc *UserController) UpdateUser(c *gin.Context) {
	var req domain.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error parsing the request"})
		return
	}

	ctx := c.Request.Context()
	err := uc.UserUseCase.UpdateUser(ctx, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// Login godoc
// @Summary      Login
// @Description  Authenticate user and return tokens
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body domain.LoginRequest true "Login credentials"
// @Success      200 {object} domain.LoginResponse
// @Failure      400 {object} map[string]string
// @Router       /public/api/users/login [post]
func (uc *UserController) Login(c *gin.Context) {
	var req domain.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error parsing the request"})
		return
	}

	ctx := c.Request.Context()
	accessToken, refreshToken, err := uc.UserUseCase.Login(ctx, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

// DeleteUser godoc
// @Summary      Delete a user
// @Description  Deletes a user by ID
// @Tags         users
// @Param        id path int true "User ID"
// @Produce      json
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Router       /private/api/users/{id} [delete]
// @Security     BearerAuth
func (uc *UserController) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscan(idParam, &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	ctx := c.Request.Context()
	if err := uc.UserUseCase.DeleteUser(ctx, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
