package controller

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/usecase"
	"github.com/gin-gonic/gin"
)

type patientController struct {
	createUseCase usecase.ICreatePatientUseCase
	deleteUseCase usecase.IDeletePatientUseCase
	updateUseCase usecase.IUpdatePatientUseCase
	findUseCase   usecase.IFindPatientsByClinicIDUseCase
}

type IPatientController interface {
	CreatePatient(c *gin.Context)
	DeletePatient(c *gin.Context)
	UpdatePatient(c *gin.Context)
	GetPatients(c *gin.Context)
}

func NewPatientController(
	createUseCase usecase.ICreatePatientUseCase,
	deleteUseCase usecase.IDeletePatientUseCase,
	updateUseCase usecase.IUpdatePatientUseCase,
	findUseCase usecase.IFindPatientsByClinicIDUseCase) IPatientController {
	return &patientController{
		createUseCase: createUseCase,
		deleteUseCase: deleteUseCase,
		updateUseCase: updateUseCase,
		findUseCase:   findUseCase,
	}
}
