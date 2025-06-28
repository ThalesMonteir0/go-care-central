package DTO

import (
	"github.com/google/uuid"
	"time"
)

type PatientDTORequest struct {
	ID             uuid.UUID `json:"id"`
	ClinicID       uuid.UUID `json:"clinic_id" binding:"required"`
	Name           string    `json:"name" binding:"required"`
	CpfResponsible string    `json:"cpf_responsible" binding:"required"`
	Cel            string    `json:"cel" binding:"required"`
	DthBirth       time.Time `json:"dth_birth" binding:"required"`
}

type PatientDTOResponse struct {
	ID             uuid.UUID `json:"id"`
	ClinicID       uuid.UUID `json:"clinic_id"`
	Name           string    `json:"name"`
	CpfResponsible string    `json:"cpf_responsible"`
	Cel            string    `json:"cel"`
	Email          string    `json:"email"`
	DthBirth       time.Time `json:"dth_birth"`
}
