package controller

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/DTO"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p patientController) CreatePatient(c *gin.Context) {
	var patientDTO DTO.PatientDTORequest
	//TODO: COLOCAR CLINIC_ID

	if err := c.ShouldBindJSON(&patientDTO); err != nil {
		c.JSON(http.StatusBadRequest, DTO.ResponseHTTP{
			Data:    nil,
			Success: false,
			Err:     "o body passado nao condiz com o esperado!",
		})
		return
	}

	if err := p.createUseCase.Execute(patientDTO); err != nil {
		c.JSON(http.StatusInternalServerError, DTO.ResponseHTTP{
			Data:    nil,
			Success: false,
			Err:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, DTO.ResponseHTTP{
		Data:    "patient criado com sucesso!",
		Success: true,
	})
}
