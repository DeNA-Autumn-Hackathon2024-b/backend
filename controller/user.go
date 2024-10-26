package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func PostUser(c echo.Context) error {
	return c.String(http.StatusOK, "Post User")
}

// e.GET("/users/:id", getUser)
func (ct *Controller) GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
