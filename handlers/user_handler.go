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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}

	for _, user := range users {
		if user.ID == id {
			return c.JSON(http.StatusOK, user)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error": "Invalid input"})
}

func CreateUser(c echo.Context) error {
	newUser := models.User{}
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	newUser.ID = len(users) + 1
	users = append(users, newUser)

	return c.JSON(http.StatusCreated, newUser)
}

func UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Id"})
	}

	for i, user := range users {
		if user.ID == id {
			if err := c.Bind(&users[i]); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})

			}
			return c.JSON(http.StatusOK, users[i])
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"error" : "User not found"})
}

func DeleteUser(c echo.Context) error {
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Id"})
	}

	for i,user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]string{"message": "User deleted"})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
}
