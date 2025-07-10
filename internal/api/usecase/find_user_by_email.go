package usecase

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/domain"
	"github.com/ThalesMonteir0/go-care-central/internal/infra/repository/postgres"
)

type findUserByEmailUseCase struct {
	repository postgres.IUserRepository
}

type IFindUserByEmailUseCase interface {
	Execute(userEmail string) (domain.User, error)
}

func NewFindUserByEmailUseCase(repository postgres.IUserRepository) IFindUserByEmailUseCase {
	return &findUserByEmailUseCase{
		repository: repository,
	}
}

func (f findUserByEmailUseCase) Execute(userEmail string) (domain.User, error) {
	user, err := f.repository.FindUserByEmail(userEmail)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
