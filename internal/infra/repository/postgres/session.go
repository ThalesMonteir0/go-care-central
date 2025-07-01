package postgres

import (
	"database/sql"
	"errors"
	"github.com/ThalesMonteir0/go-care-central/internal/api/domain"
	"github.com/google/uuid"
)

type sessionRepository struct {
	db *sql.DB
}

type ISessionRepository interface {
	Insert(session domain.Session) error
	FindSessionsByClinicID(clinicID uuid.UUID) ([]domain.Session, error)
	Delete(sessionID, clinicID uuid.UUID) error
	Update(session domain.Session) error
}

const (
	FIND_SESSIONS_QUERY = `SELECT id, dt_session, patient_id, confirmed, paid, clinic_id
							FROM sessions
							WHERE clinic_id = $1`
	INSERT_SESSION_QUERY = `INSERT INTO sessions (dt_session, patient_id, confirmed, paid, clinic_id)
								VALUES ($1, $2, $3,$4, $5)`
	UPDATE_SESSION_QUERY = `UPDATE sessions 
							SET dt_session = $1, confirmed = $2, paid = $3, report = $4
							WHERE session_id = $5 and clinic_id = $6
`
	DELETE_SESSION_QUERY = `DELETE FROM sessions WHERE session_id = $1 and clinic_id = $2`
)

func NewSessionRepository(db *sql.DB) ISessionRepository {
	return &sessionRepository{
		db: db,
	}
}

func (s sessionRepository) Insert(session domain.Session) error {
	result, err := s.db.Exec(INSERT_SESSION_QUERY,
		session.DtSession,
		session.PatientID,
		session.Confirmed,
		session.Paid,
		session.ClinicID)
	if err != nil {
		return err
	}
	rowsEffected, err := result.RowsAffected()

	if rowsEffected == 0 {
		return errors.New("nao foi possivel inserir nenhuma linha")
	}

	return nil
}

func (s sessionRepository) FindSessionsByClinicID(clinicID uuid.UUID) ([]domain.Session, error) {
	var sessions []domain.Session

	rows, err := s.db.Query(FIND_SESSIONS_QUERY, clinicID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var session domain.Session

		if err = rows.Scan(&session.ID,
			&session.DtSession,
			&session.PatientID,
			&session.Confirmed,
			&session.Paid,
			&session.ClinicID); err != nil {
			return nil, err
		}

		sessions = append(sessions, session)
	}

	return sessions, nil
}

func (s sessionRepository) Delete(sessionID, clinicID uuid.UUID) error {
	result, err := s.db.Exec(DELETE_SESSION_QUERY, sessionID, clinicID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("nao foi possivel deletar a sessão")
	}

	return nil
}

func (s sessionRepository) Update(session domain.Session) error {
	result, err := s.db.Exec(
		UPDATE_SESSION_QUERY,
		session.DtSession,
		session.Confirmed,
		session.Paid,
		session.Report,
		session.PatientID,
		session.ClinicID,
	)
	if err != nil {
		return err
	}

	rowsEffected, err := result.RowsAffected()
	if err != nil || rowsEffected == 0 {
		return errors.New("nao foi possivel atualizar sessão")
	}

	return nil
}
