package handlers

import (
	"bts-todo-app/config"
	"bts-todo-app/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Create Item
func CreateItem(c echo.Context) error {
	checklistIDstr := c.Param("id")
	checklistID, err := strconv.Atoi(checklistIDstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid checklist ID")
	}

	item := new(models.Item)
	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	item.ChecklistID = uint(checklistID)
	config.DB.Create(&item)
	return c.JSON(http.StatusCreated, item)
}

// Get Item details
func GetItemDetails(c echo.Context) error {
	id := c.Param("id")
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Item not found")
	}

	return c.JSON(http.StatusOK, item)
}

// Update Item
func UpdateItem(c echo.Context) error {
	id := c.Param("id")
	var item models.Item
	if err := json.NewDecoder(c.Request().Body).Decode(&item); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	if err := config.DB.Model(&item).Where("id = ?", id).Updates(item).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Error updating item")
	}

	return c.JSON(http.StatusOK, item)
}

// Update Item Status (Completed)
func UpdateItemStatus(c echo.Context) error {
	id := c.Param("id")
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Item not found")
	}

	item.Status = true
	if err := config.DB.Save(&item).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Error updating item status")
	}

	return c.JSON(http.StatusOK, item)
}

// Delete Item
func DeleteItem(c echo.Context) error {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Item{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Error deleting item")
	}

	return c.JSON(http.StatusOK, "Item deleted successfully")
}
