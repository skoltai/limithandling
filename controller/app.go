package controller

import (
	"errors"

	"github.com/skoltai/limithandling/domain"
	"github.com/skoltai/limithandling/store"
)

// AppController bundles the common dependencies for the controller methods
type AppController struct {
	store store.Store
}

// NewAppController constructor for AppController
func NewAppController(s store.Store) *AppController {
	return &AppController{store: s}
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

func (c *AppController) SetCustomLimits(appID int, limits []domain.Limit) {
	for _, l := range limits {
		_ = upsertLimitOverride(c.store, appID, l)
	}
}

func upsertLimitOverride(s store.Store, appID int, limit domain.Limit) error {
	limits := s.FilterLimitOverrides(func(l store.LimitOverride) bool {
		return l.AppID == appID && l.Limit.Key == limit.Key
	})

	switch len(limits) {
	case 0:
		s.CreateLimitOverride(store.LimitOverride{AppID: appID, Limit: limit})
	case 1:
		s.UpdateLimitOverride(store.LimitOverride{ID: limits[0].ID, AppID: appID, Limit: limit})
	default:
		// This should never happen
		return errors.New("Duplicated limit override")
	}

	return nil
}

func (c *AppController) OptOutPublic(appID int) {
	app, err := c.store.GetApp(appID)
	if err != nil {
		return
	}

	currentSub, err := c.store.GetSubscription(app.SubscriptionID)
	if err != nil {
		return
	}

	targetSub, ok := findSubscription(c.store, currentSub.UserID, false)
	if !ok {
		// This should never happen
		return
	}

	app.SubscriptionID = targetSub.ID
	c.store.UpdateApp(app)
}

func (c *AppController) GetLimits(appID int) ([]domain.Limit, error) {
	app, err := c.store.GetApp(appID)
	if err != nil {
		return []domain.Limit{}, err
	}

	sub, err := c.store.GetSubscription(app.SubscriptionID)
	if err != nil {
		return []domain.Limit{}, err
	}

	plan, _ := c.store.GetPlan(sub.PlanID)

	overrides := func() []domain.Limit {
		limits := c.store.FilterLimitOverrides(func(l store.LimitOverride) bool {
			return l.AppID == appID
		})
		res := make([]domain.Limit, 0)
		for _, l := range limits {
			res = append(res, l.Limit)
		}

		return res
	}()

	return domain.MergeOverrides(plan.Limits, overrides), nil
}
