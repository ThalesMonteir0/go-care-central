package controller

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/DTO"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (p patientController) GetPatients(c *gin.Context) {
	//TODO: colocar clinic_id

	var clinicID uuid.UUID

	patients, err := p.findUseCase.Execute(clinicID)
	if err != nil {
		c.JSON(http.StatusOK, DTO.ResponseHTTP{
			Data:    nil,
			Success: false,
			Err:     err.Error(),
		})
	}

	c.JSON(http.StatusOK, DTO.ResponseHTTP{
		Data:    patients,
		Success: true,
	})
}
