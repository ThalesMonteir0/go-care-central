package domain

import (
	"github.com/google/uuid"
	"time"
)

type Patient struct {
	ID             uuid.UUID
	ClinicID       uuid.UUID
	Name           string
	CpfResponsible string
	Cel            string
	DthBirth       time.Time
}
