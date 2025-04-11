package patient_useCase

import (
	"go-care-central/internal/core/domain"
	"go-care-central/pkg/rest_err"
)

type PatientUseCase struct {
	repo domain.PatientRepository
}

func NewPatientUseCase(repo domain.PatientRepository) domain.PatientUseCase {
	return PatientUseCase{
		repo: repo,
	}
}
