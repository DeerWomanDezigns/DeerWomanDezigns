package apis

import (
	"log"
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
	s.Login(c)
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
	token := s.HandleCallback(c)
	log.Println("received token", token)
	//c.JSON(http.StatusOK, token)
}
