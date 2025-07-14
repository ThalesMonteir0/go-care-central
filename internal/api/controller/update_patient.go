package controller

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/DTO"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (p patientController) UpdatePatient(c *gin.Context) {
	var patientDTO DTO.PatientDTORequest

	if err := c.ShouldBindJSON(&patientDTO); err != nil {
		c.JSON(http.StatusBadRequest, DTO.ResponseHTTP{
			Data:    nil,
			Success: false,
			Err:     "o body passado nao condiz com o esperado!",
		})
	}

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

	patientDTO.ClinicID = clinicUUID

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
