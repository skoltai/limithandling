package controller

import (
	"github.com/skoltai/limithandling/domain"
)

type AccountController struct {
	store store.Store
}

func (c *Controller) Create(user domain.User, plan domain.Plan) {
	//
}