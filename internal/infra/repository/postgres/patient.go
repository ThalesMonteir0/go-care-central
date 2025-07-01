package postgres

import (
	"database/sql"
	"github.com/ThalesMonteir0/go-care-central/internal/api/domain"
	"github.com/google/uuid"
)

type patientRepository struct {
	db *sql.DB
}

type IPatientRepository interface {
	FindPatientsByClinicID(clinicID uuid.UUID) ([]domain.Patient, error)
	Insert(patient domain.Patient) error
	Update(patient domain.Patient) error
	Delete(patientID uuid.UUID, clinicID uuid.UUID) error
}

func NewPatientRepository(DB *sql.DB) IPatientRepository {
	return &patientRepository{
		db: DB,
	}
}

func (p patientRepository) FindPatientsByClinicID(clinicID uuid.UUID) ([]domain.Patient, error) {
	var patients []domain.Patient

	query := `
			 SELECT 
			     id, 
			     name, 
			     cpf_responsible, 
			     dt_birth, 
			     cel, 
			     clinic_id 
			 FROM patient
			 WHERE clinic_id = $1
`
	rows, err := p.db.Query(query, clinicID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var patient domain.Patient

		if err = rows.Scan(
			&patient.ID,
			&patient.Name,
			&patient.CpfResponsible,
			&patient.DthBirth,
			&patient.Cel,
			&patient.ClinicID,
		); err != nil {
			return nil, err
		}

		patients = append(patients, patient)
	}

	return patients, nil
}

func (p patientRepository) Insert(patient domain.Patient) error {
	query := `
			INSERT INTO patient 
			    (name, cpf_responsible, dt_birth, cel, clinic_id) 
			values 
			    ($1, $2, $3, $4, $5)
`

	if _, err := p.db.Exec(
		query,
		patient.Name,
		patient.CpfResponsible,
		patient.DthBirth,
		patient.Cel,
		patient.ClinicID); err != nil {
		return err
	}

	return nil
}

func (p patientRepository) Update(patient domain.Patient) error {
	query := `
			UPDATE patient set 
			                   name = $1, 
			                   cpf_responsible = $2, 
			                   dt_birth = $3, 
			                   cel = $4, 
			               WHERE id = $5 AND  clinic_id= $6
			`

	if _, err := p.db.Exec(
		query,
		patient.Name,
		patient.CpfResponsible,
		patient.DthBirth,
		patient.Cel,
		patient.ID,
		patient.ClinicID); err != nil {
		return err
	}

	return nil
}

func (p patientRepository) Delete(patientID uuid.UUID, clinicID uuid.UUID) error {
	query := `
			 DELETE FROM PATIENT WHERE ID=$1 AND CLINIC_ID=$2
`

	if _, err := p.db.Exec(query, patientID, clinicID); err != nil {
		return err
	}

	return nil
}
