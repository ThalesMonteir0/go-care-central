package session_repository

import (
	"go-care-central/internal/core/domain"
	"go-care-central/internal/core/dto"
	"go-care-central/pkg/rest_err"
)

func (s SessionRepository) Create(session *dto.UserDTORequest) (*domain.Session, *rest_err.RestErr) {
	//TODO implement me
	panic("implement me")
}
