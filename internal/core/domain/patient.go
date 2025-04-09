package domain

import (
	"github.com/google/uuid"
	"go-care-central/internal/core/dto"
	"go-care-central/pkg/rest_err"
	"net/http"
	"time"
)

type Patient struct {
	ID                      uuid.UUID
	Name                    string
	CPFPersonResponsible    string
	CellOfPersonResponsible string
	DateOfBirth             time.Time
	ClinicID                int
}

type PatientService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Get(response http.ResponseWriter, request *http.Request)
	GetByID(response http.ResponseWriter, request *http.Request)
	Delete(response http.ResponseWriter, request *http.Request)
	Edit(response http.ResponseWriter, request *http.Request)
}

type PatientUseCase interface {
	Create(Patient *dto.PatientDTORequest) (*Patient, *rest_err.RestErr)
	Get() ([]Patient, *rest_err.RestErr)
	GetByID(ID string) (*Patient, *rest_err.RestErr)
	Delete(ID string) *rest_err.RestErr
	Edit(ID string) (*Patient, *rest_err.RestErr)
}

type PatientRepository interface {
	Create(Patient *dto.PatientDTORequest) (*Patient, *rest_err.RestErr)
	Get() ([]Patient, *rest_err.RestErr)
	GetByID(ID string) (*Patient, *rest_err.RestErr)
	Delete(ID string) *rest_err.RestErr
	Edit(ID string) (*Patient, *rest_err.RestErr)
}
