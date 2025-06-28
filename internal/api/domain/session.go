package domain

import (
	"github.com/google/uuid"
	"time"
)

type Session struct {
	ID        uuid.UUID
	ClinicID  uuid.UUID
	PatientID uuid.UUID
	DtSession time.Time
	Paid      bool
	Confirmed bool
	Value     float64
	Report    string
}
