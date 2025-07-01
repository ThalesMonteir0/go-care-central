package usecase

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/DTO"
	"github.com/ThalesMonteir0/go-care-central/internal/api/domain"
	"github.com/ThalesMonteir0/go-care-central/internal/infra/repository/postgres"
)

type updateUserUseCase struct {
	repository postgres.IUserRepository
}

type IUpdateUserUseCase interface {
	Execute(user DTO.UserDTORequest) error
}

func NewUpdateUserUseCase(repository postgres.IUserRepository) IUpdateUserUseCase {
	return &updateUserUseCase{
		repository: repository,
	}
}

func (u updateUserUseCase) Execute(user DTO.UserDTORequest) error {
	userDomain := domain.User{
		ID:       user.ID,
		ClinicID: user.ClinicID,
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
		Active:   user.Active,
	}

	if err := u.repository.Update(userDomain); err != nil {
		return err
	}

	return nil
}
