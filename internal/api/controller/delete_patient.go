package controller

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/DTO"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (p patientController) DeletePatient(c *gin.Context) {
	patientID := c.Param("patient_id")
	//TODO: COLOCAR CLINIC_ID

	patientIDFormated, err := uuid.Parse(patientID)
	if err != nil {
		c.JSON(http.StatusOK, DTO.ResponseHTTP{
			Data:    nil,
			Success: false,
			Err:     err.Error(),
		})
		return
	}

	if err = p.deleteUseCase.Execute(patientIDFormated, uuid.UUID{}); err != nil {
		c.JSON(http.StatusOK, DTO.ResponseHTTP{
			Data:    nil,
			Success: false,
			Err:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, DTO.ResponseHTTP{
		Data:    nil,
		Success: true,
	})
}
