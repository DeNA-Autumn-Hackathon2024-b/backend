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
	var uuidBytes [16]byte
	copy(uuidBytes[:], id)
	uuid := pgtype.UUID{
		Bytes: uuidBytes,
		Valid: true,
	}
	res, err := ct.db.GetUser(c.Request().Context(), uuid)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get user")
	}

	return c.String(http.StatusOK, res.Name)
}
