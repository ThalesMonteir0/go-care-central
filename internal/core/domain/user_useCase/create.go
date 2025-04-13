package user_useCase

import (
	"go-care-central/internal/core/domain"
	"go-care-central/internal/core/dto"
	"go-care-central/pkg/rest_err"
)

func (u UserUseCase) Create(user *dto.UserDTORequest) *rest_err.RestErr {
	userDomain := domain.User{
		ID:       user.ID,
		ClinicID: user.ClinicID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := userDomain.HashPassword(); err != nil {
		return err
	}

	user.Password = userDomain.Password
	if err := u.repo.Create(user); err != nil {
		return err
	}

	return nil
}
