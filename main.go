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
	defer i.CloseDB()
	c := controller.NewController(i)

	e.GET("/cassettes/:user_id", c.GetCassettesByUser)
	e.GET("/cassette/:cassette_id", c.GetCassette)
	e.POST("/cassette", c.CreateCassette)

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/users/:id", c.GetUser)
	e.POST("/users", c.CreateUser)

	e.GET("/cassettes/songs/:cassette_id", c.GetSongsByCassette)
	e.POST("/song", c.UploadSong)
	e.Logger.Fatal(e.Start(":8080"))
}
