package middlewares

import (
	"net/http"
	"strings"

	"github.com/arnab333/golang-app-invite-service/helpers"
	"github.com/arnab333/golang-app-invite-service/models"
	"github.com/arnab333/golang-app-invite-service/services"
	"github.com/gin-gonic/gin"
)

func VerifyToken(c *gin.Context) {
	bearToken := c.Request.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	var tokenString string
	if len(strArr) == 2 {
		tokenString = strArr[1]
	}

	claims, err := services.ExtractFromToken(tokenString, helpers.EnvKeys.JWT_ACCESS_SECRET)
	if err != nil {
		msg := "Invalid Token!"
		if strings.Contains(err.Error(), "expired") {
			msg = "Token Expired!"
		}
		c.JSON(http.StatusUnauthorized, helpers.HandleErrorResponse(msg))
		c.Abort()
		return
	}

	if claims.ID == "" {
		c.JSON(http.StatusUnauthorized, helpers.HandleErrorResponse("Invalid Token!!"))
		c.Abort()
		return
	}

	userID, err := services.FetchAuth(c, claims.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, helpers.HandleErrorResponse("Invalid Token!!!"))
		c.Abort()
		return
	}

	c.Set(helpers.CtxValues.UserID, userID)
	c.Set(helpers.CtxValues.AccessUUID, claims.ID)
	c.Set(helpers.CtxValues.Role, claims.Role)
	c.Next()
}

func VerifyAdminRole(c *gin.Context) {
	role := c.GetString(helpers.CtxValues.Role)
	if role != helpers.UserRoles.Admin {
		c.JSON(http.StatusUnauthorized, helpers.HandleErrorResponse("You are not Authorized!"))
		c.Abort()
		return
	}

	c.Next()
}

func RegisterValidation(c *gin.Context) {
	var json models.User

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, helpers.HandleErrorResponse(err.Error()))
		c.Abort()
		return
	}

	if json.FullName == "" || json.Email == "" || json.Password == "" {
		c.JSON(http.StatusBadRequest, helpers.HandleErrorResponse("Please all the fields!"))
		c.Abort()
		return
	}

	if json.Role != helpers.UserRoles.User && json.Role != helpers.UserRoles.Admin {
		c.JSON(http.StatusBadRequest, helpers.HandleErrorResponse("Roles does not match"))
		c.Abort()
		return
	}

	c.Set(helpers.CtxValues.UserDetails, &json)
	c.Next()
}
