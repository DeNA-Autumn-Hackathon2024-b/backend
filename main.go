package main

import (
	"net/http"

	"github.com/DeNA-Autumn-Hackathon2024-b/backend/controller"
	sqlc "github.com/DeNA-Autumn-Hackathon2024-b/backend/db/sqlc_gen"
	dbDriver "github.com/DeNA-Autumn-Hackathon2024-b/backend/infra"

	"github.com/DeNA-Autumn-Hackathon2024-b/backend/infra"
	"github.com/labstack/echo/v4"
	echo "github.com/labstack/echo/v4"
)

type router struct {
	echo *echo.Echo
	sql  *sqlc.Queries
}

func main() {
	e := echo.New()
	db := dbDriver.ConnectDB()
	sql := sqlc.New(db)

	router := &router{
		echo: e,
		sql:  sql,
	}

	ca := controller.NewCassette(sql)
	router.echo.GET("/cassettes", ca.GetCassette)
	router.echo.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	i := infra.NewInfrastructure()
	c := controller.NewController(i)

	router.echo.GET("/users/:id", c.GetUser)
	e.POST("/song", c.UploadSong)
	router.echo.Logger.Fatal(e.Start(":8086"))
}
