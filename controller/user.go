package controller

import (
	"net/http"

	"shape-api/form"
	"shape-api/repo"

	"github.com/gin-gonic/gin"
)

//UserController ...
type UserController struct{}

var userForm = new(form.UserForm)
var userRepo = new(repo.UserRepo)

//Register ...
func (u UserController) Register(c *gin.Context) {
	var registerForm form.RegisterForm

	if validationErr := c.ShouldBindJSON(&registerForm); validationErr != nil {
		message := userForm.Register(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	user, err := userRepo.Register(registerForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully registered", "user": user})
}

//Login
func (u UserController) Login(c *gin.Context) {
	var loginForm form.LoginForm

	if validationErr := c.ShouldBindJSON(&loginForm); validationErr != nil {
		message := userForm.Register(validationErr)
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": message})
		return
	}

	user, token, err := userRepo.Login(loginForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully login", "user": user, "token": token})
}
