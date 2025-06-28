package domain

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	ClinicID uuid.UUID
	Name     string
	Password string
	Email    string
	Active   bool
}
