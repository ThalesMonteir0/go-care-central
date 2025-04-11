package session_useCase

import (
	"go-care-central/internal/core/domain"
)

type SessionUseCase struct {
	repo domain.SessionRepository
}

func NewSessionUseCase(repo domain.SessionRepository) domain.SessionUseCase {
	return SessionUseCase{
		repo: repo,
	}
}
