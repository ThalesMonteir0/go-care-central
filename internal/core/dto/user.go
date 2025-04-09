package dto

import (
	"encoding/json"
	"go-care-central/pkg/rest_err"
	"io"
)

type UserDTORequest struct {
	ID       int    `json:"id"`
	ClinicID int    `json:"clinic_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func FromJsonUserDTORequest(body io.Reader) (*UserDTORequest, *rest_err.RestErr) {
	var UserDTO UserDTORequest
	if err := json.NewDecoder(body).Decode(&UserDTO); err != nil {
		return nil, rest_err.NewBadRequestError("")
	}

	return &UserDTO, nil
}
