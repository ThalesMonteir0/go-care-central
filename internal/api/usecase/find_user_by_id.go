package usecase

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/DTO"
	"github.com/ThalesMonteir0/go-care-central/internal/api/repository/postgres"
	"github.com/google/uuid"
)

type findUserByIDUseCase struct {
	repository postgres.IUserRepository
}

type IFindUserByIDUseCase interface {
	Execute(userID uuid.UUID) (DTO.UserDTOResponse, error)
}

func NewFindUserByIDUseCase(repository postgres.IUserRepository) IFindUserByIDUseCase {
	return &findUserByIDUseCase{
		repository: repository,
	}
}

func (f findUserByIDUseCase) Execute(userID uuid.UUID) (DTO.UserDTOResponse, error) {
	user, err := f.repository.FindUserByID(userID)
	if err != nil {
		return DTO.UserDTOResponse{}, err
	}

	userDto := DTO.UserDTOResponse{
		ID:       user.ID,
		ClinicID: user.ClinicID,
		Name:     user.Name,
		Email:    user.Email,
		Active:   user.Active,
	}

	return userDto, nil
}
