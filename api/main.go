package main

import (
	"fmt"
	"log"

	"github.com/deer-woman-dezigns/deer-woman-dezigns/code/apis"
	"github.com/deer-woman-dezigns/deer-woman-dezigns/code/config"
	"github.com/deer-woman-dezigns/deer-woman-dezigns/code/middleware"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/guregu/dynamo"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/deer-woman-dezigns/deer-woman-dezigns/docs"
)

// @title Deer Woman Dezigns Swagger API
// @version 1.0
// @description Swagger API for Deer Woman Dezigns.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email deerwomandezigns.site@gmail.com

// @license.name Creative Commons
// @license.url https://github.com/Maybeenaught/DeerWomanDezigns/blob/main/license.md

// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := config.LoadConfig("config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

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
		v1.Use(middleware.CORS())
		v1.Use(middleware.Auth())
		v1.GET("/users", apis.GetAllUsers)
		v1.GET("/users/:id", apis.GetUser)
		v1.POST("/users", apis.AddUser)
		v1.PUT("/users/:id", apis.ModifyUser)
		v1.DELETE("/users/:id", apis.DeleteUser)
		v1.GET("/etsy/login", apis.EtsyLogin)
		v1.GET("/etsy/callback", apis.EtsyCallback)
	}

	sess := session.Must(session.NewSession())
	config.Config.DB = dynamo.New(sess, &aws.Config{Region: aws.String(`us-east-2`)})

	log.Println("Successfully connected to database")

	r.Run(fmt.Sprintf(":%v", config.Config.ServerPort))
}
