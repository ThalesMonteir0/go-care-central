package usecase

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/DTO"
	"github.com/ThalesMonteir0/go-care-central/internal/api/repository/postgres"
	"github.com/google/uuid"
)

type findSessionsByClinicIDUseCase struct {
	repository postgres.ISessionRepository
}

type IFindSessionsByClinicIDUseCase interface {
	Execute(clinicID uuid.UUID) ([]DTO.SessionDTO, error)
}

func NewFindSessionsByClinicIDUseCase(repository postgres.ISessionRepository) IFindSessionsByClinicIDUseCase {
	return &findSessionsByClinicIDUseCase{
		repository: repository,
	}
}

func (f findSessionsByClinicIDUseCase) Execute(clinicID uuid.UUID) ([]DTO.SessionDTO, error) {
	var sessionsDTO []DTO.SessionDTO

	sessions, err := f.repository.FindSessionsByClinicID(clinicID)
	if err != nil {
		return nil, err
	}

	for _, session := range sessions {
		sessionDto := DTO.SessionDTO{
			ID:        session.ID,
			ClinicID:  session.ClinicID,
			PatientID: session.PatientID,
			DtSession: session.DtSession,
			Paid:      session.Paid,
			Value:     session.Value,
			Confirmed: session.Confirmed,
			Report:    session.Report,
		}

		sessionsDTO = append(sessionsDTO, sessionDto)
	}

	return sessionsDTO, nil
}
