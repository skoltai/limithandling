package controller

import (
	"github.com/skoltai/limithandling/domain"
	"github.com/skoltai/limithandling/store"
)

type AccountController struct {
	store store.Store
}

func (c *AccountController) Create(user domain.User, plan domain.Plan) {
	//
}
