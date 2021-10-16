package apis

import (
	"net/http"

	"github.com/deer-woman-dezigns/deer-woman-dezigns/code/services"
	"github.com/gin-gonic/gin"
)

// EtsyLogin godoc
// @Summary Initiates authentication to Etsy
// @Produce json
// @Success 200
// @Router /etsy/login [get]
// @Security ApiKeyAuth
func EtsyLogin(c *gin.Context) {
	s := services.NewEtsyService()
	csrfToken := s.RandState(32)
	maxAge := 12 * 60 * 60
	c.SetCookie("csrfCookie", csrfToken, maxAge, "/", "deerwoman-dezigns.com", true, false)
	s.Login(csrfToken)
	c.JSON(http.StatusOK, "")
}

// EtsyCallback godoc
// @Summary Handles callback from Etsy
// @Produce json
// @Success 200
// @Router /etsy/callback [get]
// @Security ApiKeyAuth
func EtsyCallback(c *gin.Context) {
	s := services.NewEtsyService()
	state := c.Query("state")
	code := c.Query("code")
	if csrfToken, err := c.Cookie("csrfCookie"); err != nil || csrfToken != state {
		c.JSON(http.StatusUnauthorized, "CSRF token invalid or expired")
	} else {
		s.HandleCallback(code)
		c.JSON(http.StatusOK, "")
	}
}
