package domain

import (
	"go-care-central/internal/core/dto"
	"go-care-central/pkg/rest_err"
	"net/http"
)

type Clinic struct {
	ID         int
	Name       string
	Cnpj       string
	EmailOwner string
	CelOwner   string
}

type ClinicService interface {
	Create(response http.ResponseWriter, request *http.Request)
}

type ClinicUseCase interface {
	Create(clinic *dto.ClinicDTOrequest) (*dto.ClinicDTOrequest, *rest_err.RestErr)
}

type ClinicRepository interface {
	Create(clinic *dto.ClinicDTOrequest) (*Clinic, *rest_err.RestErr)
}
