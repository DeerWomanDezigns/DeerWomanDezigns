package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Maybeenaught/deer-woman-dezigns/code/apis"
	"github.com/Maybeenaught/deer-woman-dezigns/code/config"
	"github.com/Maybeenaught/deer-woman-dezigns/code/httputil"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Maybeenaught/deer-woman-dezigns/code/docs"
)

// @title Deer Woman Dezigns Swagger API
// @version 1.0
// @description Swagger API for Deer Woman Dezigns.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email deerwomandezigns.site@gmail.com

// @license.name MIT
// @license.url https://github.com/Maybeenaught/DeerWomanDezigns/blob/main/license.md

// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// load application configurations
	if err := config.LoadConfig("config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.Use(auth())
		v1.GET("/users/:id", apis.GetUser)
	}

	config.Config.DB, config.Config.DBErr = gorm.Open("postgres", config.Config.DSN)
	if config.Config.DBErr != nil {
		panic(config.Config.DBErr)
	}

	// config.Config.DB.AutoMigrate(&models.User{}) // This is needed for generation of schema for postgres image.

	defer config.Config.DB.Close()

	log.Println("Successfully connected to database")

	r.Run(fmt.Sprintf(":%v", config.Config.ServerPort))
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			httputil.NewError(c, http.StatusUnauthorized, errors.New("authorization is a required header"))
			c.Abort()
		}
		if authHeader != config.Config.ApiKey {
			httputil.NewError(c, http.StatusUnauthorized, fmt.Errorf("this user isn't authorized to this operation: api_key=%s", authHeader))
			c.Abort()
		}
		c.Next()
	}
}
