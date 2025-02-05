package routes

import (
	"bts-todo-app/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/register", handlers.RegisterUser)
	e.POST("/login", handlers.LoginUser)

	e.POST("/checklist", handlers.CreateChecklist)
	e.GET("/checklist", handlers.GetChecklists)
	e.GET("/checklist/:id", handlers.GetChecklistDetails)
	e.DELETE("/checklist/:id", handlers.DeleteChecklist)

	e.POST("/checklists/:id/items", handlers.CreateItem)
	e.GET("/items/:id", handlers.GetItemDetails)
	e.PUT("/items/:id", handlers.UpdateItem)
	e.PATCH("/items/:id/status", handlers.UpdateItemStatus)
	e.DELETE("/items/:id", handlers.DeleteItem)
}
