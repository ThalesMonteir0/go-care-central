package controller

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/DTO"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (n sessionController) CreateSession(c *gin.Context) {
	var session DTO.SessionDTO

	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, DTO.ResponseHTTP{
			Data:    nil,
			Success: false,
			Err:     err.Error(),
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

	session.ClinicID = clinicUUID

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
