package user_repository

import (
	"database/sql"
	"go-care-central/internal/core/domain"
	"go-care-central/internal/core/dto"
	"go-care-central/pkg/rest_err"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return UserRepository{
		db: db,
	}
}
