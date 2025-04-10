package session_repository

import (
	"database/sql"
	"go-care-central/internal/core/domain"
)

type SessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) domain.SessionRepository {
	return SessionRepository{
		db: db,
	}
}
