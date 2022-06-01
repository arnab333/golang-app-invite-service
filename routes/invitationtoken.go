package routes

import (
	"github.com/arnab333/golang-app-invite-service/controllers"
	"github.com/arnab333/golang-app-invite-service/middlewares"
	"github.com/gin-gonic/gin"
)

func invitationTokenRoutes(rg *gin.RouterGroup) {
	rg.GET("/verify-invitation/:inviteToken", middlewares.Throttle(1, 3), controllers.VerifyInvitationToken)

	rg.GET("/invitation-tokens", middlewares.VerifyToken, middlewares.VerifyAdminRole, controllers.GetAllInvitationTokens)

	rg.POST("/create-invitation-token", middlewares.VerifyToken, middlewares.VerifyAdminRole, controllers.CreateInvitationToken)

	rg.DELETE("/disable-invitation-token", middlewares.VerifyToken, middlewares.VerifyAdminRole, controllers.DisableInvitationToken)
}
