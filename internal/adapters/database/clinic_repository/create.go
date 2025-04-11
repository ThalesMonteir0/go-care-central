package clinic_repository

import (
	"go-care-central/internal/core/domain"
	"go-care-central/internal/core/dto"
	"go-care-central/pkg/rest_err"
)

func (c ClinicRepository) Create(clinic *dto.ClinicDTORequest) (*domain.Clinic, *rest_err.RestErr) {
	//TODO implement me
	panic("implement me")
}
