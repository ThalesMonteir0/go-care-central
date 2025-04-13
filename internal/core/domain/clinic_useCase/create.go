package clinic_useCase

import (
	"go-care-central/internal/core/dto"
	"go-care-central/pkg/rest_err"
)

func (c ClinicUseCase) Create(clinic *dto.ClinicDTORequest) *rest_err.RestErr {
	err := c.repository.Create(clinic)
	if err != nil {
		return err
	}

	return nil
}
