package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/arnab333/golang-app-invite-service/helpers"
	"github.com/arnab333/golang-app-invite-service/models"
	"github.com/arnab333/golang-app-invite-service/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	user := c.Keys[helpers.CtxValues.UserDetails].(*models.User)

	user.IsActive = true

	result, err := services.DBConn.FindUser(&models.User{Email: user.Email})

	if err != nil {
		fmt.Println("FindUser err ==>", err.Error())
	}

	if result.ID != 0 {
		c.JSON(http.StatusConflict, helpers.HandleErrorResponse("Email Already Exists!"))
		c.Abort()
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		fmt.Println("GenerateFromPassword err ==>", err.Error())
		c.JSON(http.StatusUnauthorized, helpers.HandleErrorResponse(err.Error()))
		c.Abort()
		return
	}

	user.Password = string(password)

	if _, err := services.DBConn.InsertUser(user); err != nil {
		fmt.Println("InsertUser err ==>", err)
		c.JSON(http.StatusBadRequest, helpers.HandleErrorResponse(err.Error()))
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, helpers.HandleSuccessResponse(helpers.CreatedMessage, nil))
}

func Login(c *gin.Context) {
	var json map[string]string

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, helpers.HandleErrorResponse(err.Error()))
		c.Abort()
		return
	}

	user, err := services.DBConn.FindUser(&models.User{Email: json["email"]})
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.HandleErrorResponse("Invalid Email or Password!!"))
		c.Abort()
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(json["password"])); err != nil {
		c.JSON(http.StatusNotFound, helpers.HandleErrorResponse("Invalid Email or Password!"))
		c.Abort()
		return
	}

	td, err := services.CreateAuth(c, strconv.FormatUint(uint64(user.ID), 10), user.Role)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		c.Abort()
		return
	}

	tokens := gin.H{
		"accessToken":  td.AccessToken,
		"refreshToken": td.RefreshToken,
	}
	c.JSON(http.StatusCreated, helpers.HandleSuccessResponse("", tokens))
}

func Logout(c *gin.Context) {
	accessUUID := c.GetString(helpers.CtxValues.AccessUUID)
	if accessUUID == "" {
		c.JSON(http.StatusUnauthorized, helpers.HandleErrorResponse("Unauthorized!"))
		c.Abort()
		return
	}

	deleted, err := services.DeleteAuth(c, accessUUID)
	if err != nil || deleted == 0 {
		c.JSON(http.StatusUnauthorized, helpers.HandleErrorResponse("Unauthorized!!"))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, helpers.HandleSuccessResponse("Successfully logged out", nil))
}

func RefreshToken(c *gin.Context) {
	var json map[string]string

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusUnprocessableEntity, helpers.HandleErrorResponse(err.Error()))
		c.Abort()
		return
	}

	tokenString := json["refreshToken"]

	claims, err := services.ExtractFromToken(tokenString, helpers.EnvKeys.JWT_REFRESH_SECRET)
	if err != nil {
		msg := "Invalid Token!"
		if strings.Contains(err.Error(), "expired") {
			msg = "Token Expired!"
		}
		c.JSON(http.StatusUnauthorized, helpers.HandleErrorResponse(msg))
		c.Abort()
		return
	}

	deleted, err := services.DeleteAuth(c, claims.ID)
	if err != nil || deleted == 0 {
		c.JSON(http.StatusUnauthorized, helpers.HandleErrorResponse("Unauthorized!!"))
		c.Abort()
		return
	}

	td, err := services.CreateAuth(c, claims.UserID, claims.Role)
	if err != nil {
		c.JSON(http.StatusForbidden, err.Error())
		c.Abort()
		return
	}

	tokens := gin.H{
		"accessToken":  td.AccessToken,
		"refreshToken": td.RefreshToken,
	}
	c.JSON(http.StatusCreated, helpers.HandleSuccessResponse("", tokens))
}
