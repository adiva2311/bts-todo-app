package routes

import (
	"bts-todo-app/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/register", handlers.RegisterUser)
	e.POST("/login", handlers.LoginUser)

	e.POST("/checklist", handlers.CreateChecklist)
	e.POST("/update_checklist/:id", handlers.UpdateChecklists)
	e.GET("/checklist", handlers.GetChecklists)
	e.GET("/checklist/:id", handlers.GetChecklistDetails)
	e.DELETE("/checklist/:id", handlers.DeleteChecklist)

	e.POST("/checklists/:id/items", handlers.CreateItem) // ID = id checklist
	e.GET("/items/:id", handlers.GetItemDetails) // ID = id item
	e.PUT("/items/:id", handlers.UpdateItem) // ID = id item
	e.PATCH("/items/:id/status", handlers.UpdateItemStatus) // ID = id item
	e.DELETE("/items/:id", handlers.DeleteItem) // ID = id item
}
