package apis

import (
	"log"
	"net/http"
	"strconv"

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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "Malformed request body"})
	} else if _, err = strconv.Atoi(newUser.ID); err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "ID must be valid int"})
	} else if user, err := s.Add(newUser); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// ModifyUser godoc
// @Summary Modifies and returns user based on given id and json
// @Accept json
// @Param id path integer true "User ID"
// @Param data body models.User true "body data"
// @Produce json
// @Success 200 {object} models.User
// @Router /users/{id} [put]
// @Security ApiKeyAuth
func ModifyUser(c *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO())
	id := c.Param("id")

	var newUserValues models.User
	if c.ShouldBind(&newUserValues) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "Malformed request body"})
		log.Println("Malformed request body")
	} else if newUserValues.ID != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "ID should not be given in request body for PUT"})
	} else {
		var user *models.User
		var err error
		if user, err = s.Modify(id, newUserValues); err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			log.Println(err)
		} else {
			c.JSON(http.StatusOK, &user)
		}
	}
}
