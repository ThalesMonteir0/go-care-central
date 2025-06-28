package usecase

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/DTO"
	"github.com/ThalesMonteir0/go-care-central/internal/api/repository/postgres"
)

type findUserByEmailUseCase struct {
	repository postgres.IUserRepository
}

type IFindUserByEmailUseCase interface {
	Execute(userEmail string) (DTO.UserLoginDTOResponse, error)
}

func NewFindUserByEmailUseCase(repository postgres.IUserRepository) IFindUserByEmailUseCase {
	return &findUserByEmailUseCase{
		repository: repository,
	}
}

func (f findUserByEmailUseCase) Execute(userEmail string) (DTO.UserLoginDTOResponse, error) {
	user, err := f.repository.FindUserByEmail(userEmail)
	if err != nil {
		return DTO.UserLoginDTOResponse{}, err
	}

	userDto := DTO.UserLoginDTOResponse{
		ID:       user.ID,
		ClinicID: user.ClinicID,
		Name:     user.Name,
		Email:    user.Email,
		Active:   user.Active,
	}

	return userDto, nil
}
