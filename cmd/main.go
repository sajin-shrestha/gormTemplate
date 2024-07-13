package main

import (
	"net/http"

	"github.com/sajin-shrestha/gormTemplate/config"
	"github.com/sajin-shrestha/gormTemplate/controllers"
	"github.com/sajin-shrestha/gormTemplate/models"
	"github.com/sajin-shrestha/gormTemplate/routes"
)

func main() {
	config.ConnectDB()

	models.Migrate(config.DB)

	controllers.Initialize(config.DB)

	r := routes.InitializeRoutes()
	http.ListenAndServe(":8000", r)
}
