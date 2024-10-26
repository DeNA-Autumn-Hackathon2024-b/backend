package main

import (
	"net/http"

	"github.com/DeNA-Autumn-Hackathon2024-b/backend/controller"

	"github.com/DeNA-Autumn-Hackathon2024-b/backend/infra"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	i := infra.NewInfrastructure()
	c := controller.NewController(i)
	e.GET("/cassettes", c.GetCassettesByUser)
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/users/:id", c.GetUser)
	e.POST("/song", c.UploadSong)
	e.Logger.Fatal(e.Start(":8080"))
}
