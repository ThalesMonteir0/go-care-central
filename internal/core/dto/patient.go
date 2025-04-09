package dto

import (
	"encoding/json"
	"github.com/google/uuid"
	"go-care-central/pkg/rest_err"
	"io"
	"time"
)

type PatientDTORequest struct {
	ID                      uuid.UUID `json:"id"`
	Name                    string    `json:"name"`
	CPFPersonResponsible    string    `json:"cpf_person_responsible"`
	CellOfPersonResponsible string    `json:"cell_of_person_responsible"`
	DateOfBirth             time.Time `json:"date_of_birth"`
	ClinicID                int       `json:"clinic_id"`
}

func FromJSONPatientRequest(body io.Reader) (*PatientDTORequest, *rest_err.RestErr) {
	var PatientDomain PatientDTORequest
	if err := json.NewDecoder(body).Decode(&PatientDomain); err != nil {
		return nil, rest_err.NewBadRequestError("")
	}

	return &PatientDomain, nil
}
