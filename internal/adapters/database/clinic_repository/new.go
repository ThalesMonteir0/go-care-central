package clinic_repository

import (
	"database/sql"
	"go-care-central/internal/core/domain"
)

type ClinicRepository struct {
	db *sql.DB
}

func NewClinicRepository(db *sql.DB) domain.ClinicRepository {
	return ClinicRepository{
		db: db,
	}
}
