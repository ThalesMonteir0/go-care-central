package usecase

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/repository/postgres"
	"github.com/google/uuid"
)

type deletePatientUseCase struct {
	repository postgres.IPatientRepository
}

type IDeletePatientUseCase interface {
	Execute(patientID uuid.UUID, clinicID uuid.UUID) error
}

func NewDeletePatientUseCase(repository postgres.IPatientRepository) IDeletePatientUseCase {
	return &deletePatientUseCase{
		repository: repository,
	}
}

func (d deletePatientUseCase) Execute(patientID uuid.UUID, clinicID uuid.UUID) error {
	if err := d.repository.Delete(patientID, clinicID); err != nil {
		return err
	}

	return nil
}
