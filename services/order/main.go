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
			"message": "Welcome to the Order Service!",
		})
	})

	e.POST("/orders", createOrder)

	// Start the server
	e.Logger.Fatal(e.Start(":8082"))
}

// Handler to create a new order
func createOrder(c echo.Context) error {
	type Order struct {
		ProductID int     `json:"product_id"`
		UserID    int     `json:"user_id"`
		Quantity  int     `json:"quantity"`
		Total     float64 `json:"total"`
	}

	order := new(Order)
	if err := c.Bind(order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	// Simulate order total calculation
	order.Total = float64(order.Quantity) * 100 // Assume each product costs 100

	return c.JSON(http.StatusCreated, order)
}
