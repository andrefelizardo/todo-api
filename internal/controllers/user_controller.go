package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{
	// CreateUser func(c *gin.Context)
}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created",
	})
}