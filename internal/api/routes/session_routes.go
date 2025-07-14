package routes

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/controller"
	"github.com/gin-gonic/gin"
)

func InitSessionRoutes(r *gin.RouterGroup, controller controller.SessionsController) {
	r.GET("/", controller.GetSessions)
	r.POST("/", controller.CreateSession)
	r.PUT("/", controller.UpdateSession)
	r.DELETE("/:session_id", controller.DeleteSession)
}
