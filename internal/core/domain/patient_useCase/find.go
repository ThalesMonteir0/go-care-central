package patient_useCase

import (
	"go-care-central/internal/core/domain"
	"go-care-central/pkg/rest_err"
)

func (p PatientUseCase) Get() ([]domain.Patient, *rest_err.RestErr) {
	//TODO implement me
	panic("implement me")
}

func (p PatientUseCase) GetByID(ID string) (*domain.Patient, *rest_err.RestErr) {
	//TODO implement me
	panic("implement me")
}
