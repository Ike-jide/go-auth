package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/ike-jide/go-auth/controllers"
	"github.com/ike-jide/go-auth/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("users", controller.GetUsers())
	incomingRoutes.GET("users/:user_id", controller.Getuser())
}
