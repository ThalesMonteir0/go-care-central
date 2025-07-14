package routes

import (
	"github.com/ThalesMonteir0/go-care-central/internal/api/controller"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.Engine, controller controller.UserController) {
	r.POST("/login", controller.Login)
}
