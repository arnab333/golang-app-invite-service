package controllers

import (
	"net/http"

	"github.com/arnab333/golang-app-invite-service/helpers"
	"github.com/arnab333/golang-app-invite-service/services"
	"github.com/gin-gonic/gin"
)

func VerifyInvitationToken(c *gin.Context) {
	inviteToken := c.Param("inviteToken")

	if inviteToken == "" {
		c.JSON(http.StatusBadRequest, helpers.HandleErrorResponse(helpers.RequiredInvitationToken))
		c.Abort()
		return
	}

	_, err := services.GetInviteToken(inviteToken)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.HandleErrorResponse("Invalid Invitation Token!"))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, helpers.HandleSuccessResponse("Login Success!", nil))
}

func CreateInvitationToken(c *gin.Context) {
	userID := c.GetString(helpers.CtxValues.UserID)
	it, err := services.CreateInviteToken(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.HandleErrorResponse("Sorry! Some Error Occurred!"))
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, helpers.HandleSuccessResponse("", helpers.Map{
		"token": it,
	}))
}

func DisableInvitationToken(c *gin.Context) {
	var json helpers.Map

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, helpers.HandleErrorResponse("Invalid Invitation Token!"))
		c.Abort()
		return
	}

	if json["token"] == "" {
		c.JSON(http.StatusBadRequest, helpers.HandleErrorResponse(helpers.RequiredInvitationToken))
		c.Abort()
		return
	}

	result := services.DisableInviteToken(json["token"].(string))

	c.JSON(http.StatusOK, helpers.HandleSuccessResponse("", result))
}

func GetAllInvitationTokens(c *gin.Context) {
	result := services.GetAllInviteTokens()
	c.JSON(http.StatusOK, helpers.HandleSuccessResponse("", result))
}
