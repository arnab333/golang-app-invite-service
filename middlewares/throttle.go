package middlewares

import (
	"errors"
	"net/http"

	"github.com/arnab333/golang-app-invite-service/helpers"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func Throttle(maxEventsPerSec int, maxBurstSize int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(maxEventsPerSec), maxBurstSize)

	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.Error(errors.New("limit exceeded"))
			c.JSON(http.StatusUnauthorized, helpers.HandleErrorResponse("Too Many Requests!"))
			c.Abort()
			return
		}

		c.Next()
	}
}
