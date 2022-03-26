package routes

import (
	"github.com/Wizhumanz/golang-jwt-project/controllers"
	"github.com/Wizhumanz/golang-jwt-project/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/user/:userId", controllers.GetUser())
}
