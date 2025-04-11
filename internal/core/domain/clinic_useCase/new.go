package clinic_useCase

import (
	"go-care-central/internal/core/domain"
	"go-care-central/internal/core/dto"
	"go-care-central/pkg/rest_err"
)

type ClinicUseCase struct {
	repository domain.ClinicRepository
}

func NewClinicUseCase(repo domain.ClinicRepository) domain.ClinicUseCase {
	return ClinicUseCase{
		repository: repo,
	}
}
