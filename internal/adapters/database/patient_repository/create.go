package patient_repository

import (
	"go-care-central/internal/core/domain"
	"go-care-central/internal/core/dto"
	"go-care-central/pkg/rest_err"
)

func (p PatientRepository) Create(Patient *dto.PatientDTORequest) (*domain.Patient, *rest_err.RestErr) {
	//TODO implement me
	panic("implement me")
}
