package controller

import "github.com/DeNA-Autumn-Hackathon2024-b/backend/infra"

type Controller struct {
	infra *infra.Infrastructure
}

func NewController(infra *infra.Infrastructure) *Controller {
	return &Controller{
		infra: infra,
	}
}
