package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(cors.Default())

	api := router.Group("/api")
	v1 := api.Group("/v1")

	auth := v1.Group("/")
	invitationToken := v1.Group("/")

	authRoutes(auth)
	invitationTokenRoutes(invitationToken)

	return router
}
