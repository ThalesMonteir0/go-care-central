package user_service

import (
	"go-care-central/internal/core/domain"
)

type UserService struct {
	useCase domain.UserUseCase
}

func NewUserService(useCase domain.UserUseCase) domain.UserService {
	return UserService{
		useCase: useCase,
	}
}
