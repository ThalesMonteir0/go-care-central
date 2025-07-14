package routes

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/controller"
	"github.com/gin-gonic/gin"
)

func InitPatientRoutes(r *gin.RouterGroup, controller controller.IPatientController) {
	r.GET("/", controller.GetPatients)
	r.POST("/", controller.CreatePatient)
	r.DELETE("/:patient_id", controller.DeletePatient)
	r.PUT("/", controller.UpdatePatient)
}
