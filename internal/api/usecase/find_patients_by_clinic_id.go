package usecase

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/DTO"
	"github.com/ThalesMonteir0/go-care-central/internal/infra/repository/postgres"
	"github.com/google/uuid"
)

type findPatientsByClinicIDUseCase struct {
	repository postgres.IPatientRepository
}

type IFindPatientsByClinicIDUseCase interface {
	Execute(clinicID uuid.UUID) ([]DTO.PatientDTOResponse, error)
}

func NewFindPatientsByClinicIDUseCase(repository postgres.IPatientRepository) IFindPatientsByClinicIDUseCase {
	return &findPatientsByClinicIDUseCase{
		repository: repository,
	}
}

func (f findPatientsByClinicIDUseCase) Execute(clinicID uuid.UUID) ([]DTO.PatientDTOResponse, error) {
	var patientsDTO []DTO.PatientDTOResponse

	patients, err := f.repository.FindPatientsByClinicID(clinicID)
	if err != nil {
		return nil, err
	}

	for _, patient := range patients {
		patientDTO := DTO.PatientDTOResponse{
			ID:             patient.ID,
			ClinicID:       patient.ClinicID,
			Name:           patient.Name,
			CpfResponsible: patient.CpfResponsible,
			Cel:            patient.Cel,
			DthBirth:       patient.DthBirth,
		}

		patientsDTO = append(patientsDTO, patientDTO)
	}

	return patientsDTO, nil
}
