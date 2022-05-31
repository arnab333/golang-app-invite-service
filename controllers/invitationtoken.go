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
		c.JSON(http.StatusUnprocessableEntity, helpers.HandleErrorResponse(helpers.RequiredInvitationToken))
		c.Abort()
		return
	}
}

func CreateInvitationToken(c *gin.Context) {
	userID := c.GetString(helpers.CtxValues.UserID)
	it, err := services.CreateInviteToken(c, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.HandleErrorResponse("Sorry! Some Error Occurred!"))
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, helpers.HandleSuccessResponse("", helpers.Map{
		"inviteToken": it,
	}))
}

func DeleteInvitationToken(c *gin.Context) {
	var json helpers.Map

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, helpers.HandleErrorResponse(err.Error()))
		c.Abort()
		return
	}

	if json["inviteToken"] == "" {
		c.JSON(http.StatusBadRequest, helpers.HandleErrorResponse(helpers.RequiredInvitationToken))
		c.Abort()
		return
	}

	deleted, err := services.DeleteAuth(c, json["inviteToken"].(string))
	if err != nil || deleted == 0 {
		c.JSON(http.StatusUnauthorized, helpers.HandleErrorResponse("Unauthorized!!"))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, helpers.HandleSuccessResponse("Successfully deleted token", nil))
}
