package usecase

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/repository/postgres"
	"github.com/google/uuid"
)

type deleteSessionsUseCase struct {
	repository postgres.ISessionRepository
}

type IDeleteSessionsUseCase interface {
	Execute(sessionID uuid.UUID, clinicID uuid.UUID) error
}

func NewDeleteSessionsUseCase(repository postgres.ISessionRepository) IDeleteSessionsUseCase {
	return &deleteSessionsUseCase{
		repository: repository,
	}
}

func (f deleteSessionsUseCase) Execute(sessionID, clinicID uuid.UUID) error {
	if err := f.repository.Delete(sessionID, clinicID); err != nil {
		return err
	}

	return nil
}
