package dto

import (
	"encoding/json"
	"go-care-central/pkg/rest_err"
	"io"
)

type ClinicDTORequest struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Cnpj       string `json:"cnpj"`
	EmailOwner string `json:"email_owner"`
	CelOwner   string `json:"cel_owner"`
}

func FromJsonClinicRequest(body io.Reader) (*ClinicDTORequest, *rest_err.RestErr) {
	var clinicDTO ClinicDTORequest
	if err := json.NewDecoder(body).Decode(&clinicDTO); err != nil {
		return nil, rest_err.NewBadRequestError("")
	}

	return &clinicDTO, nil
}
