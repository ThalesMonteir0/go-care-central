package patient_repository

import (
	"database/sql"
	"go-care-central/internal/core/domain"
	"go-care-central/pkg/rest_err"
)

type PatientRepository struct {
	db *sql.DB
}

func NewPatientRepository(db *sql.DB) domain.PatientRepository {
	return PatientRepository{
		db: db,
	}
}
