package main

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/controller"
	"github.com/ThalesMonteir0/go-care-central/internal/api/routes"
	"github.com/ThalesMonteir0/go-care-central/internal/api/usecase"
	"github.com/ThalesMonteir0/go-care-central/internal/config"
	"github.com/ThalesMonteir0/go-care-central/internal/infra/repository/postgres"
	"github.com/gin-gonic/gin"
)

func main() {
	env := config.LoadVariables()
	r := gin.Default()
	userController, sessionController, patientController := initDependecies(*env)

	rPatient := r.Group("/patient")
	routes.InitPatientRoutes(rPatient, patientController)
	rSession := r.Group("/session")
	routes.InitSessionRoutes(rSession, sessionController)
	routes.InitUserRoutes(r, userController)

	r.Run(":3001")
}

func initDependecies(env config.Env) (controller.UserController, controller.SessionsController, controller.IPatientController) {
	db := postgres.NewPostgresConnection(env)
	patientRepo := postgres.NewPatientRepository(db)
	sessionRepo := postgres.NewSessionRepository(db)
	userRepo := postgres.NewUserRepository(db)

	createPatientUC := usecase.NewCreatePatientUseCase(patientRepo)
	deletePatientUC := usecase.NewDeletePatientUseCase(patientRepo)
	updatePatientUC := usecase.NewUpdatePatientUseCase(patientRepo)
	findPatientUC := usecase.NewFindPatientsByClinicIDUseCase(patientRepo)

	createSessionUC := usecase.NewCreateSessionsUseCase(sessionRepo)
	deleteSessionUC := usecase.NewDeleteSessionsUseCase(sessionRepo)
	updateSessionUC := usecase.NewUpdateSessionsUseCase(sessionRepo)
	findSessionsUC := usecase.NewFindSessionsByClinicIDUseCase(sessionRepo)

	findUserUC := usecase.NewFindUserByEmailUseCase(userRepo)
	loginUC := usecase.NewLoginUseCase(findUserUC)

	loginController := controller.NewUserController(loginUC)
	sessionController := controller.NewSessionController(
		deleteSessionUC,
		createSessionUC,
		findSessionsUC,
		updateSessionUC,
	)
	patientController := controller.NewPatientController(
		createPatientUC,
		deletePatientUC,
		updatePatientUC,
		findPatientUC,
	)

	return loginController, sessionController, patientController
}
