package apis

import (
	"log"
	"net/http"

	"github.com/deer-woman-dezigns/deer-woman-dezigns/code/daos"
	"github.com/deer-woman-dezigns/deer-woman-dezigns/code/models"
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
	id := c.Param("id")
	if user, err := s.Get(id); err != nil {
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

// AddUser godoc
// @Summary Adds and returns user based on given json
// @Accept json
// @Param data body models.User true "body data"
// @Produce json
// @Success 200 {object} models.User
// @Router /users [post]
// @Security ApiKeyAuth
func AddUser(c *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO())
	var err error
	var newUser models.User
	if c.ShouldBind(&newUser) != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		log.Println(err)
	}
	if user, err := s.Add(newUser); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
