package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EtsyTest(c *gin.Context) {
	c.AbortWithStatus(http.StatusUnauthorized)
}
