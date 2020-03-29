package controller

import (
	"github.com/skoltai/limithandling/domain"
	"github.com/skoltai/limithandling/store"
)

type AppController struct {
	store store.Store
}

func (c *AppController) Create(user domain.User, app domain.App) {
	//
}

func (c *AppController) SetCustomLimits(app domain.App, limits []domain.Limit) {
	//
}

func (c *AppController) OptOutPublic(app domain.App) {
	//
}

func (c *AppController) GetLimits(app domain.App) {
	//
}
