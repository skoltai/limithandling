package controller

import (
	"github.com/skoltai/limithandling/domain"
	"github.com/skoltai/limithandling/store"
)

type AppController struct {
	store store.Store
}

func (c *Controller) Create(user domain.User, app domain.App) {
	//
}

func (c *Controller) SetCustomLimits(app domain.App, limits []domain.Limit) {
	//
}

func (c *Controller) OptOutPublic(app domain.App) {
	//
}

func (c *Controller) GetLimits(app domain.App) {
	//
}
