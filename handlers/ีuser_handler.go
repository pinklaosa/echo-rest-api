package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yourusername/echo-rest-api/models"
)

var users = []models.User{
	{ID: 1, Name: "John Doe", Email: "john@example.com"},
	{ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
}

func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World !!")
}

func GetUser(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func GetUserByID(c echo.Context) error {
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest , map[string]string{"error": "Invalid ID"})
	}

	for _,user := range users {
		if user.ID == id {
			return c.JSON(http.StatusOK, user)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "Invalid input"})
}


