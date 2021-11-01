package apis

import (
	"net/http"

	"github.com/deer-woman-dezigns/deer-woman-dezigns/code/services"
	"github.com/gin-gonic/gin"
)

// GetTest godoc
// @Summary Returns a 401
// @Produce json
// @Success 200
// @Router /api/v1/etsy/test [get]
// @Security ApiKeyAuth
func EtsyTest(c *gin.Context) {
	if _, err := c.Request.Cookie("access_token"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "etsy tokens are missing and need to be acquired"})
	} else {
		c.JSON(http.StatusOK, "Cookies set")
	}
}

func SetTokens(c *gin.Context) {
	params := c.Request.URL.Query()
	redirectUri := params.Get("redirect_uri")
	code := params.Get("code")
	codeVer := params.Get("code_verifier")

	s := services.NewEtsyService()
	s.GetAuthToken(c, code, codeVer, redirectUri)
	c.JSON(http.StatusOK, "Tokens acquired")
}
