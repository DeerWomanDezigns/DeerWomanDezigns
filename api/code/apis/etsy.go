package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTest godoc
// @Summary Returns a 401
// @Produce json
// @Success 200
// @Router /api/v1/etsy/test [get]
// @Security ApiKeyAuth
func EtsyTest(c *gin.Context) {
	c.AbortWithStatus(http.StatusUnauthorized)
}
