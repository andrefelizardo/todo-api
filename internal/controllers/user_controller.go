package controllers

import (
	"fmt"
	"net/http"

	"github.com/andrefelizardo/todo-api/internal/request"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct{
	// CreateUser func(c *gin.Context)
}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) CreateUser(ctx *gin.Context) {

	var input request.CreateUserRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Validation error: %s", errors)})
		return
	}
	
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created",
	})
}