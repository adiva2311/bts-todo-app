package handlers

import (
	"bts-todo-app/config"
	"bts-todo-app/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	result := config.DB.Create(user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to register user"})
	}

	return c.JSON(http.StatusOK, user)
}

func LoginUser(c echo.Context) error {
	var user models.User
	input := new(models.User)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	config.DB.Where("username = ?", input.Username).First(&user)
	if user.ID == 0 || user.Password != input.Password {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Login successful"})
}
