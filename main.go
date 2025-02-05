package main

import (
	"bts-todo-app/config"
	"bts-todo-app/models"
	"bts-todo-app/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.ConnectDatabase()

	// Migrate models
	config.DB.AutoMigrate(&models.User{}, &models.Checklist{}, &models.Item{})

	// API Routing
	routes.InitRoutes(e)

	e.Start(":5000")
}
