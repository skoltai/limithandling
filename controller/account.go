package controller

import (
	"github.com/skoltai/limithandling/domain"
	"github.com/skoltai/limithandling/store"
)

// AccountController bundles the common dependencies for the controller methods
type AccountController struct {
	ur store.UserRepository
	sr store.SubscriptionRepository
}

// NewAccountController constructor for AccountController
func NewAccountController(ur store.UserRepository, sr store.SubscriptionRepository) *AccountController {
	return &AccountController{
		ur: ur,
		sr: sr,
	}
}

// Create creates a user record and associates a subscription with it
func (c *AccountController) Create(user domain.User, planID int) {
	userID := c.ur.Create(user)
	s := store.Subscription{
		UserID: userID,
		PlanID: planID,
		Subscription: domain.Subscription{
			Public: false,
		},
	}
	c.sr.Create(s)
}
