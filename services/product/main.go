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
			"message": "Welcome to the Product Service!",
		})
	})

	e.GET("/products", getProducts)
	e.POST("/products", createProduct)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handeler to fetch all products
func getProducts(c echo.Context) error {
	//Dummy product data
	products := []map[string]interface{}{
		{"id": 1, "name": "Product A", "price": 100},
		{"id": 2, "name": "Product B", "price": 200},
	}
	return c.JSON(http.StatusOK, products)
}

// Handler t crate a new product
func createProduct(c echo.Context) error {
	type Product struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}

	product := new(Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid input",
		})
	}

	return c.JSON(http.StatusCreated, product)
}
