package controller

import (
	"net/http"

	sqlc "github.com/DeNA-Autumn-Hackathon2024-b/backend/db/sqlc_gen"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type Cassette struct {
	db *sqlc.Queries
}

func NewCassette(db *sqlc.Queries) *Cassette {
	return &Cassette{db: db}
}

func (c *Controller) PostCassette(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Post Cassette")
}

func (c *Controller) GetCassettesByUser(ctx echo.Context) error {
	userID := ctx.Param("user_id")
	var uuid pgtype.UUID
	err := uuid.Scan(userID)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid UUID format")
	}
	cassettes, err := c.db.GetCassettesByUser(ctx.Request().Context(), uuid)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to get cassettes")
	}
	return ctx.JSON(http.StatusOK, cassettes)
}

func (c *Controller) GetCassette(ctx echo.Context) error {
	userID := ctx.Param("user_id")
	var uuid pgtype.UUID
	err := uuid.Scan(userID)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid UUID format")
	}
	res, err := c.db.GetCassette(ctx.Request().Context(), uuid)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to get cassette")
	}
	return ctx.JSON(http.StatusOK, res)
}
