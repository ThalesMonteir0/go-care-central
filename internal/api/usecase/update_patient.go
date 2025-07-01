package usecase

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/DTO"
	"github.com/ThalesMonteir0/go-care-central/internal/api/domain"
	"github.com/ThalesMonteir0/go-care-central/internal/infra/repository/postgres"
)

type updatePatientUseCase struct {
	repository postgres.IPatientRepository
}

type IUpdatePatientUseCase interface {
	Execute(patient DTO.PatientDTORequest) error
}

func NewUpdatePatientUseCase(repository postgres.IPatientRepository) IUpdatePatientUseCase {
	return &updatePatientUseCase{
		repository: repository,
	}
}

func (u updatePatientUseCase) Execute(patient DTO.PatientDTORequest) error {
	patientDomain := domain.Patient{
		ID:             patient.ID,
		ClinicID:       patient.ClinicID,
		Name:           patient.Name,
		CpfResponsible: patient.CpfResponsible,
		Cel:            patient.Cel,
		DthBirth:       patient.DthBirth,
	}

	if err := u.repository.Update(patientDomain); err != nil {
		return err
	}

	return nil
}
