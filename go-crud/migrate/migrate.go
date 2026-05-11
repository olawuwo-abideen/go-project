package main

import (
	"github.com/olawuwo-abideen/go-crud/initializers"
	"github.com/olawuwo-abideen/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
