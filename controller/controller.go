package controller

import (
	sqlc "github.com/DeNA-Autumn-Hackathon2024-b/backend/db/sqlc_gen"
	"github.com/DeNA-Autumn-Hackathon2024-b/backend/infra"
)

type Controller struct {
	infra *infra.Infrastructure
	db    *sqlc.Queries
}

func NewController(infra *infra.Infrastructure) *Controller {
	sql, err := infra.ConnectDB()
	if err != nil {
		panic(err)
	}
	return &Controller{
		infra: infra,
		db:    sqlc.New(sql),
	}
}
