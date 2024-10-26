package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/DeNA-Autumn-Hackathon2024-b/backend/controller"
)

func main() {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/users/:id", controller.GetUser)
	e.Logger.Fatal(e.Start(":8080"))
}
