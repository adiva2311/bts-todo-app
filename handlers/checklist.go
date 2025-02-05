package handlers

import (
	"bts-todo-app/config"
	"bts-todo-app/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Create Checklist
func CreateChecklist(c echo.Context) error {
	//var checklist models.Checklist
	// if err := json.NewDecoder(c.Request().Body).Decode(&checklist); err != nil {
	// 	return c.JSON(http.StatusBadRequest, "Invalid input")
	// }

	checklist := new(models.Checklist)
	if err := c.Bind(checklist); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	result := config.DB.Create(&checklist)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add checklist"})
	}
	return c.JSON(http.StatusCreated, checklist)
}

// Get all checklists
func GetChecklists(c echo.Context) error {
	var checklists []models.Checklist
	if err := config.DB.Find(&checklists).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Error fetching checklists")
	}

	return c.JSON(http.StatusOK, checklists)
}

// Get Checklist details
func GetChecklistDetails(c echo.Context) error {
	id := c.Param("id")
	var checklist models.Checklist
	if err := config.DB.Preload("Items").First(&checklist, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Checklist not found")
	}

	return c.JSON(http.StatusOK, checklist)
}

// Delete Checklist
func DeleteChecklist(c echo.Context) error {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Checklist{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Error deleting checklist")
	}

	return c.JSON(http.StatusOK, "Checklist deleted successfully")
}
