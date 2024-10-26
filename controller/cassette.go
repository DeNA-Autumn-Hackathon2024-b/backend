package controller

import (
	"net/http"

	sqlc "github.com/DeNA-Autumn-Hackathon2024-b/backend/db/sqlc_gen"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type Cassette struct {
	db *sqlc.Queries
}

func NewCassette(db *sqlc.Queries) *Cassette {
	return &Cassette{db: db}
}

func (ca *Cassette) PostCassette(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Post Cassette")
}

func (ca *Cassette) GetCassettesByUser(ctx echo.Context) error {
	userID := ctx.Param("user_id")
	cassettes, err := ca.db.GetCassettesByUser(ctx.Request().Context(), pgtype.UUID{UUID: userID})
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Failed to get cassettes")
	}
	return ctx.JSON(http.StatusOK, cassettes)
}

func (ca *Cassette) GetCassette(ctx echo.Context) error {
	userID := ctx.Param("user_id")
	uuid, err := uuid.Parse(userID)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "Invalid UUID")
	}
	ca.db.GetCassette(ctx.Request().Context(), uuid)
	return ctx.String(http.StatusOK, id)
}
