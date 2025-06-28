package controller

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/DTO"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (n newSessionController) CreateSession(c *gin.Context) {
	var session DTO.SessionDTO
	//todo: clinic_id do token

	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, DTO.ResponseHTTP{
			Data:    nil,
			Success: false,
			Err:     err.Error(),
		})
	}

	if err := n.createSessionUseCase.Execute(session); err != nil {
		c.JSON(http.StatusInternalServerError, DTO.ResponseHTTP{
			Data:    nil,
			Success: false,
			Err:     err.Error(),
		})
	}

	c.JSON(http.StatusCreated, DTO.ResponseHTTP{
		Data:    "success",
		Success: true,
		Err:     "",
	})
}
