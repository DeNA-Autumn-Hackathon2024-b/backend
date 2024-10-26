package controller

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

func PostUser(c echo.Context) error {
	return c.String(http.StatusOK, "Post User")
}

// e.GET("/users/:id", getUser)
func (ct *Controller) GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	var uuid pgtype.UUID
	err := uuid.Scan(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid UUID format")
	}
	res, err := ct.db.GetUser(c.Request().Context(), uuid)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get user")
	}

	return c.JSON(http.StatusOK, res)
}
