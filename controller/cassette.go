package controller

import (
	"net/http"

	sqlc "github.com/DeNA-Autumn-Hackathon2024-b/backend/db/sqlc_gen"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type CreateCassetteRequest struct {
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}

func (c *Controller) CreateCassette(ctx echo.Context) error {
	var req CreateCassetteRequest
	if err := ctx.Bind(&req); err != nil {
		return err
	}
	var uuid pgtype.UUID
	err := uuid.Scan(req.UserID)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid UUID format")
	}
	res, err := c.db.PostCassette(ctx.Request().Context(), sqlc.PostCassetteParams{
		Name:   req.Name,
		UserID: uuid,
	})
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to create cassette")
	}
	return ctx.String(http.StatusOK, res.Name)
}

func (c *Controller) GetCassettesByUser(ctx echo.Context) error {
	userID := ctx.Param("cassette_id")
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
