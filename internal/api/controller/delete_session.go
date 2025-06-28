package controller

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/DTO"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (n newSessionController) DeleteSession(c *gin.Context) {
	sessionID := c.Param("session_id")
	SessionUUID, err := uuid.Parse(sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, DTO.ResponseHTTP{
			Data:    nil,
			Success: false,
			Err:     err.Error(),
		})
		return
	}
	clinicID := uuid.UUID{}
	//todo: pegar o clinic_id do token

	if err = n.deleteSessionUseCase.Execute(SessionUUID, clinicID); err != nil {
		c.JSON(http.StatusInternalServerError, DTO.ResponseHTTP{
			Data:    nil,
			Success: false,
			Err:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, DTO.ResponseHTTP{
		Data:    "success",
		Success: true,
		Err:     "",
	})
}
