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
	c.JSON(http.StatusUnauthorized, gin.H{"error": "etsy tokens are missing and need to be acquired"})
}
