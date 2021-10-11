package apis

import (
	"log"
	"net/http"
	"strconv"

	"github.com/deer-woman-dezigns/deer-woman-dezigns/code/daos"
	"github.com/deer-woman-dezigns/deer-woman-dezigns/code/services"
	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
// @Security ApiKeyAuth
func GetUser(c *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO())
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	if user, err := s.Get(int(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// GetAllUsers godoc
// @Summary Retrieves all users
// @Produce json
// @Success 200 {object} []models.User
// @Router /users [get]
// @Security ApiKeyAuth
func GetAllUsers(c *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO())
	if users, err := s.GetAll(); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, users)
	}
}
