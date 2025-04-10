package domain

import (
	"go-care-central/internal/core/dto"
	"go-care-central/pkg/rest_err"
	"net/http"
	"time"
)

type Session struct {
	ID           string
	PatientID    string
	DateWithHour time.Time
	Confirmed    bool
	Paid         bool
	Record       string
}

type SessionService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Update(response http.ResponseWriter, request *http.Request)
	Delete(response http.ResponseWriter, request *http.Request)
	Paid(response http.ResponseWriter, request *http.Request)
	ConfirmByWpp(response http.ResponseWriter, request *http.Request)
}

type SessionUseCase interface {
	Create(session *dto.UserDTORequest) (*dto.UserDTORequest, *rest_err.RestErr)
	Update(session *dto.UserDTORequest) (*dto.UserDTORequest, *rest_err.RestErr)
	Delete(sessionID string) *rest_err.RestErr
	Paid(sessionID string) *rest_err.RestErr
	ConfirmByWpp()
}

type SessionRepository interface {
	Create(session *dto.UserDTORequest) (*Session, *rest_err.RestErr)
	Update(session *dto.UserDTORequest) (*Session, *rest_err.RestErr)
	Delete(sessionID string) *rest_err.RestErr
	Paid(sessionID string) *rest_err.RestErr
}
