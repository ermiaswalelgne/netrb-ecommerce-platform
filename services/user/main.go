package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize Echo
	e := echo.New()

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Welcome to the User Service!",
		})
	})

	e.POST("/users", createUser)

	// Start the server
	e.Logger.Fatal(e.Start(":8081"))
}

// Handler to create a new user
func createUser(c echo.Context) error {
	type User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	user := new(User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	return c.JSON(http.StatusCreated, user)
}
