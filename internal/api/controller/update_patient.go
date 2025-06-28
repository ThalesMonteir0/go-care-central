package controller

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/DTO"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p patientController) UpdatePatient(c *gin.Context) {
	var patientDTO DTO.PatientDTORequest
	//TODO: COLOCAR CLINIC_ID

	if err := c.ShouldBindJSON(&patientDTO); err != nil {
		c.JSON(http.StatusBadRequest, DTO.ResponseHTTP{
			Data:    nil,
			Success: false,
			Err:     "o body passado nao condiz com o esperado!",
		})
	}

	if err := p.updateUseCase.Execute(patientDTO); err != nil {
		c.JSON(http.StatusBadRequest, DTO.ResponseHTTP{
			Data:    nil,
			Success: false,
			Err:     err.Error(),
		})
	}

	c.JSON(http.StatusOK, DTO.ResponseHTTP{
		Data:    "atualizado com sucesso!",
		Success: true,
		Err:     "",
	})
}
