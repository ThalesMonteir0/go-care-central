package usecase

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/DTO"
	"github.com/ThalesMonteir0/go-care-central/internal/api/domain"
	"github.com/ThalesMonteir0/go-care-central/internal/api/repository/postgres"
)

type createSessionsUseCase struct {
	repository postgres.ISessionRepository
}

type ICreateSessionsUseCase interface {
	Execute(session DTO.SessionDTO) error
}

func NewCreateSessionsUseCase(repository postgres.ISessionRepository) ICreateSessionsUseCase {
	return &createSessionsUseCase{
		repository: repository,
	}
}

func (f createSessionsUseCase) Execute(session DTO.SessionDTO) error {
	sessionDomain := domain.Session{
		ClinicID:  session.ClinicID,
		PatientID: session.PatientID,
		DtSession: session.DtSession,
		Paid:      session.Paid,
		Confirmed: session.Confirmed,
		Value:     session.Value,
		Report:    session.Report,
	}

	if err := f.repository.Insert(sessionDomain); err != nil {
		return err
	}

	return nil
}
