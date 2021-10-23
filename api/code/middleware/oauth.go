package middleware

import (
	"net/http"

	"github.com/deer-woman-dezigns/deer-woman-dezigns/code/services"
	"github.com/gin-gonic/gin"
)

func EtsyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := c.Cookie("access_token"); err != nil {
			if refreshToken, err := c.Cookie("refresh_token"); err != nil {
				c.Redirect(http.StatusTemporaryRedirect, "/login")
			} else {
				s := services.NewEtsyService()
				s.RefreshAuthToken(c, refreshToken)
			}
		}

		c.Next()
	}
}
