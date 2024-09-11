package controllers

import (
	"net/http"

	"github.com/andrefelizardo/todo-api/internal/request"
	"github.com/andrefelizardo/todo-api/internal/usecases"
	"github.com/gin-gonic/gin"
)

type UserController struct{
	userUsecase usecases.UserUseCase
}

func NewUserController(userUsecase usecases.UserUseCase) *UserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

func (u *UserController) CreateUser(ctx *gin.Context) {

	var input request.CreateUserRequest

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.userUsecase.CreateUser(input)
	if err != nil {
		if validationErr, ok := err.(*usecases.ValidationError); ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Message})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created",
		"data": user,
	})
}