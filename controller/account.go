package controller

import (
	"github.com/skoltai/limithandling/domain"
	"github.com/skoltai/limithandling/store"
)

// AccountController bundles the common dependencies for the controller methods
type AccountController struct {
	store store.Store
}

// Create creates a user record and associates a subscription with it
func (c *AccountController) Create(user domain.User, planID int) {
	userID := c.store.AddUser(user)
	s := store.Subscription{
		UserID: userID,
		PlanID: planID,
		Subscription: domain.Subscription{
			Public: false,
		},
	}
	c.store.CreateSubscription(s)
}
