package main

import (
	"github.com/dortlii/simple-api/initializers"
	"github.com/dortlii/simple-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
