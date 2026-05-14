package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/olawuwo-abideen/go-project/go-auth/controllers"
	"github.com/olawuwo-abideen/go-project/go-auth/initializers"

	_ "github.com/olawuwo-abideen/go-project/go-auth/docs"
)

// @title Go CRUD API
// @version 1.0
// @description Simple CRUD API using Gin and GORM
// @host localhost:3000
// @BasePath /

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {

	r := gin.Default()

	// Routes
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run server
	r.Run()
}
