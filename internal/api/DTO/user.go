package DTO

import "github.com/google/uuid"

type UserDTORequest struct {
	ID       uuid.UUID `json:"id"`
	ClinicID uuid.UUID `json:"clinic_id"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	Active   bool      `json:"active"`
}

type UserDTOResponse struct {
	ID       uuid.UUID `json:"id"`
	ClinicID uuid.UUID `json:"clinic_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Active   bool      `json:"active"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginDTOResponse struct {
	ID       uuid.UUID `json:"id"`
	ClinicID uuid.UUID `json:"clinic_id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Active   bool      `json:"active"`
}
