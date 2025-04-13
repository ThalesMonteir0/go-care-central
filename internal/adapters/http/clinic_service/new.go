package clinic_service

import (
	"go-care-central/internal/core/domain"
)

type ClinicService struct {
	useCase domain.ClinicUseCase
}

func NewClinicService(useCase domain.ClinicUseCase) domain.ClinicService {
	return ClinicService{
		useCase: useCase,
	}
}
