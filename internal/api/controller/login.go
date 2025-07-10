package controller

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/DTO"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u userController) Login(c *gin.Context) {
	var userLoginDTO DTO.UserLoginRequest

	if err := c.ShouldBindJSON(&userLoginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid json"})
		return
	}

	jwt, err := u.loginUseCase.Execute(userLoginDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "unable to login user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": jwt,
	})
}
