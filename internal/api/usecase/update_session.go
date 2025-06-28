package usecase

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/DTO"
	"github.com/ThalesMonteir0/go-care-central/internal/api/domain"
	"github.com/ThalesMonteir0/go-care-central/internal/api/repository/postgres"
)

type updateSessionsUseCase struct {
	repository postgres.ISessionRepository
}

type IUpdateSessionsUseCase interface {
	Execute(session DTO.SessionDTO) error
}

func NewUpdateSessionsUseCase(repository postgres.ISessionRepository) IUpdateSessionsUseCase {
	return &updateSessionsUseCase{
		repository: repository,
	}
}

func (f updateSessionsUseCase) Execute(session DTO.SessionDTO) error {
	sessionDomain := domain.Session{
		ClinicID:  session.ClinicID,
		PatientID: session.PatientID,
		DtSession: session.DtSession,
		Paid:      session.Paid,
		Value:     session.Value,
		Confirmed: session.Confirmed,
		Report:    session.Report,
	}

	if err := f.repository.Update(sessionDomain); err != nil {
		return err
	}

	return nil
}
