package controller

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/usecase"
	"github.com/gin-gonic/gin"
)

type newSessionController struct {
	deleteSessionUseCase usecase.IDeleteSessionsUseCase
	updateSessionUseCase usecase.IUpdateSessionsUseCase
	findSessionsUseCase  usecase.IFindSessionsByClinicIDUseCase
	createSessionUseCase usecase.ICreateSessionsUseCase
}

type NewSessionsController interface {
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
) NewSessionsController {
	return &newSessionController{
		deleteSessionUseCase: deleteUseCase,
		updateSessionUseCase: UpdateSessionUseCase,
		findSessionsUseCase:  GetSessionUseCase,
		createSessionUseCase: CreateSessionUseCase,
	}
}
