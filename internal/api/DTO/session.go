package DTO

import (
	"github.com/google/uuid"
	"time"
)

type SessionDTO struct {
	ID        uuid.UUID `json:"id"`
	ClinicID  uuid.UUID `json:"clinic_id" binding:"required"`
	PatientID uuid.UUID `json:"patient_id" binding:"required"`
	DtSession time.Time `json:"dt_session" binding:"required"`
	Paid      bool      `json:"paid" binding:"required"`
	Confirmed bool      `json:"confirmed" binding:"required"`
	Value     float64   `json:"value" binding:"required"`
	Report    string    `json:"report" binding:"required"`
}
