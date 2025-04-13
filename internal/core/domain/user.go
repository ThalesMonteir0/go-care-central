package domain

import (
	"go-care-central/internal/core/dto"
	"go-care-central/pkg/rest_err"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type User struct {
	ID       int
	ClinicID int
	Name     string
	Email    string
	Password string
}

func (u User) HashPassword() *rest_err.RestErr {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return rest_err.NewInternalServeError("Could not encrypt password")
	}

	u.Password = string(hash)

	return nil
}

type UserService interface {
	Create(response http.ResponseWriter, request *http.Request)
}

type UserUseCase interface {
	Create(user *dto.UserDTORequest) *rest_err.RestErr
}

type UserRepository interface {
	Create(user *dto.UserDTORequest) *rest_err.RestErr
}
