package controller

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/DTO"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (n sessionController) GetSessions(c *gin.Context) {
	clinicID, ok := c.Get("clinic_id")
	if !ok {
		c.JSON(http.StatusBadRequest, DTO.ResponseHTTP{
			Data:    nil,
			Success: false,
			Err:     "nao foi possivel buscar o clinic_id deste usuario",
		})
		return
	}

	clinicUUID, err := uuid.Parse(clinicID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, DTO.ResponseHTTP{
			Data:    nil,
			Success: false,
			Err:     "nao foi possivel converter o clinic_id",
		})
		return
	}

	sessions, err := n.findSessionsUseCase.Execute(clinicUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, DTO.ResponseHTTP{
			Data:    nil,
			Success: false,
			Err:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, DTO.ResponseHTTP{
		Data:    sessions,
		Success: true,
		Err:     "",
	})
}
