package usecase

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/DTO"
	"github.com/ThalesMonteir0/go-care-central/internal/api/domain"
	"github.com/ThalesMonteir0/go-care-central/internal/infra/repository/postgres"
)

type createPatientUseCase struct {
	repository postgres.IPatientRepository
}

type ICreatePatientUseCase interface {
	Execute(patient DTO.PatientDTORequest) error
}

func NewCreatePatientUseCase(repository postgres.IPatientRepository) ICreatePatientUseCase {
	return &createPatientUseCase{
		repository: repository,
	}
}

func (c createPatientUseCase) Execute(patient DTO.PatientDTORequest) error {
	patientDomain := domain.Patient{
		ID:             patient.ID,
		ClinicID:       patient.ClinicID,
		Name:           patient.Name,
		CpfResponsible: patient.CpfResponsible,
		Cel:            patient.Cel,
		DthBirth:       patient.DthBirth,
	}

	if err := c.repository.Insert(patientDomain); err != nil {
		return err
	}

	return nil
}
