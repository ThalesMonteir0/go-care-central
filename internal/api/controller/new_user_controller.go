package controller

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/usecase"
	"github.com/gin-gonic/gin"
)

type userController struct {
	loginUseCase usecase.LoginUseCase
}

type UserController interface {
	Login(c *gin.Context)
}

func NewUserController(loginUseCase usecase.LoginUseCase) UserController {
	return &userController{
		loginUseCase: loginUseCase,
	}
}
