package domain

import (
	"go-care-central/internal/core/dto"
	"go-care-central/pkg/rest_err"
	"net/http"
)

type User struct {
	ID       int
	ClinicID int
	Name     string
	Email    string
	Password string
}

type UserService interface {
	Create(response http.ResponseWriter, request *http.Request)
}

type UserUseCase interface {
	Create(user *dto.UserDTORequest) (*dto.UserDTORequest, *rest_err.RestErr)
}

type UserRepository interface {
	Create(user *dto.UserDTORequest) (*Clinic, *rest_err.RestErr)
}
