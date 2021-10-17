package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/deer-woman-dezigns/deer-woman-dezigns/code/config"
	"github.com/deer-woman-dezigns/deer-woman-dezigns/code/httputil"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.FullPath() == "/api/v1/etsy/callback" {
			c.Next()
			return
		}
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			httputil.NewError(c, http.StatusUnauthorized, errors.New("authorization is a required header"))
			c.Abort()
		}
		if authHeader != config.Config.ApiKey {
			httputil.NewError(c, http.StatusUnauthorized, fmt.Errorf("this user isn't authorized to this operation: api_key=%s", authHeader))
			c.Abort()
		}
		c.Next()
	}
}
