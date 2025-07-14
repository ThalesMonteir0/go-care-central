package controller

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/usecase"
	"github.com/gin-gonic/gin"
)

type sessionController struct {
	deleteSessionUseCase usecase.IDeleteSessionsUseCase
	updateSessionUseCase usecase.IUpdateSessionsUseCase
	findSessionsUseCase  usecase.IFindSessionsByClinicIDUseCase
	createSessionUseCase usecase.ICreateSessionsUseCase
}

type SessionsController interface {
	CreateSession(c *gin.Context)
	DeleteSession(c *gin.Context)
	UpdateSession(c *gin.Context)
	GetSessions(c *gin.Context)
}

func NewSessionController(
	deleteUseCase usecase.IDeleteSessionsUseCase,
	UpdateSessionUseCase usecase.IUpdateSessionsUseCase,
	GetSessionUseCase usecase.IFindSessionsByClinicIDUseCase,
	CreateSessionUseCase usecase.ICreateSessionsUseCase,
) SessionsController {
	return &sessionController{
		deleteSessionUseCase: deleteUseCase,
		updateSessionUseCase: UpdateSessionUseCase,
		findSessionsUseCase:  GetSessionUseCase,
		createSessionUseCase: CreateSessionUseCase,
	}
}
