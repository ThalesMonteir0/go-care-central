package dto

import (
	"encoding/json"
	"github.com/google/uuid"
	"go-care-central/pkg/rest_err"
	"io"
	"time"
)

type SessionDTORequest struct {
	ID           uuid.UUID `json:"id"`
	PatientID    uuid.UUID `json:"patient_id"`
	DateWithHour time.Time `json:"date_with_hour"`
	Confirmed    bool      `json:"confirmed"`
	Paid         bool      `json:"paid"`
	Record       string    `json:"record"`
}

func FromJsonSessionDTORequest(body io.Reader) (*SessionDTORequest, *rest_err.RestErr) {
	var SessionDTO SessionDTORequest
	if err := json.NewDecoder(body).Decode(&SessionDTO); err != nil {
		return nil, rest_err.NewBadRequestError("")
	}

	return &SessionDTO, nil
}
