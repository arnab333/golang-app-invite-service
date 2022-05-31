package routes

import (
	"github.com/arnab333/golang-app-invite-service/controllers"
	"github.com/arnab333/golang-app-invite-service/middlewares"
	"github.com/gin-gonic/gin"
)

func invitationTokenRoutes(rg *gin.RouterGroup) {
	rg.GET("/verify-invite/:inviteToken", middlewares.Throttle(1, 3), controllers.VerifyInvitationToken)

	rg.POST("/create-invitation-token", middlewares.VerifyToken, middlewares.VerifyAdminRole, controllers.CreateInvitationToken)

	rg.DELETE("/delete-invitation-token", middlewares.VerifyToken, middlewares.VerifyAdminRole, controllers.DeleteInvitationToken)
}
