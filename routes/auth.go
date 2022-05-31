package routes

import (
	"github.com/arnab333/golang-app-invite-service/controllers"
	"github.com/arnab333/golang-app-invite-service/middlewares"
	"github.com/gin-gonic/gin"
)

func authRoutes(rg *gin.RouterGroup) {
	rg.POST("/register", middlewares.RegisterValidation, controllers.Register)

	rg.POST("/login", controllers.Login)

	rg.POST("/logout", middlewares.VerifyToken, controllers.Logout)

	rg.POST("/refresh-token", middlewares.VerifyToken, controllers.RefreshToken)

}
