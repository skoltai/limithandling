package controller

import (
	"errors"

	"github.com/skoltai/limithandling/domain"
	"github.com/skoltai/limithandling/store"
)

type AppController struct {
	store store.Store
}

func (c *AppController) Create(userID int, app domain.App) {
	subID, err := func() (int, error) {
		// Creating private apps are more common, so checking that first
		if app.Public == false {
			if s, ok := findSubscription(c.store, userID, false); ok {
				return s.ID, nil
			}

			// This should never happen
			return 0, errors.New("No private plan for user")
		}

		if sub, ok := findSubscription(c.store, userID, true); ok {
			return sub.ID, nil
		}

		return createPublicSubscription(c.store, userID), nil
	}()

	if err != nil {
		return
	}

	c.store.CreateApp(store.App{
		OwnerID:        userID,
		SubscriptionID: subID,
		App:            app,
	})
}

func findSubscription(s store.Store, userID int, private bool) (store.Subscription, bool) {
	return s.FindSubscription(func(sub store.Subscription) bool {
		return sub.UserID == userID && sub.Subscription.Public == private
	})
}

func createPublicSubscription(s store.Store, userID int) int {
	return s.CreateSubscription(store.Subscription{
		UserID: userID,
		PlanID: 4, // bad magic number
		Subscription: domain.Subscription{
			Public: true,
		},
	})
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
