package usecase

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/DTO"
	"github.com/ThalesMonteir0/go-care-central/internal/api/domain"
	"github.com/ThalesMonteir0/go-care-central/internal/infra/auth"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type loginUsecase struct {
	findByEmail IFindUserByEmailUseCase
}

type LoginUseCase interface {
	Execute(request DTO.UserLoginRequest) (string, error)
}

func NewLoginUseCase(findByEmailUseCase IFindUserByEmailUseCase) LoginUseCase {
	return &loginUsecase{
		findByEmail: findByEmailUseCase,
	}
}

func (l loginUsecase) Execute(request DTO.UserLoginRequest) (string, error) {
	user, err := l.findByEmail.Execute(request.Email)
	if err != nil {
		return "", err
	}

	if user.ID == uuid.Nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return "", err
	}

	userDTO := createResponseDTO(user)

	jwt, err := auth.CreateJWT(userDTO)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func createResponseDTO(user domain.User) DTO.UserDTOResponse {
	return DTO.UserDTOResponse{
		ID:       user.ID,
		ClinicID: user.ClinicID,
		Name:     user.Name,
		Email:    user.Email,
		Active:   user.Active,
	}
}
