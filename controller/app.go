package controller

import (
	"errors"

	"github.com/skoltai/limithandling/domain"
	"github.com/skoltai/limithandling/store"
)

// AppController bundles the common dependencies for the controller methods
type AppController struct {
	sr  store.SubscriptionRepository
	ar  store.AppRepository
	lor store.LimitOverrideRepository
	pr  store.PlanRepository
}

// NewAppController constructor for AppController
func NewAppController(sr store.SubscriptionRepository, ar store.AppRepository, lor store.LimitOverrideRepository, pr store.PlanRepository) *AppController {
	return &AppController{
		sr:  sr,
		ar:  ar,
		lor: lor,
		pr:  pr,
	}
}

func (c *AppController) Create(userID int, app domain.App) {
	subID, err := func() (int, error) {
		// Creating private apps are more common, so checking that first
		if app.Public == false {
			if s, ok := findSubscription(c.sr, userID, false); ok {
				return s.ID, nil
			}

			// This should never happen
			return 0, errors.New("No private plan for user")
		}

		if sub, ok := findSubscription(c.sr, userID, true); ok {
			return sub.ID, nil
		}

		return createPublicSubscription(c.sr, userID), nil
	}()

	if err != nil {
		return
	}

	c.ar.Create(store.App{
		OwnerID:        userID,
		SubscriptionID: subID,
		App:            app,
	})
}

func findSubscription(sr store.SubscriptionRepository, userID int, private bool) (store.Subscription, bool) {
	return sr.Find(func(sub store.Subscription) bool {
		return sub.UserID == userID && sub.Subscription.Public == private
	})
}

func createPublicSubscription(sr store.SubscriptionRepository, userID int) int {
	return sr.Create(store.Subscription{
		UserID: userID,
		PlanID: 4, // bad magic number
		Subscription: domain.Subscription{
			Public: true,
		},
	})
}

func (c *AppController) SetCustomLimits(appID int, limits []domain.Limit) {
	for _, l := range limits {
		_ = upsertLimitOverride(c.lor, appID, l)
	}
}

func upsertLimitOverride(lor store.LimitOverrideRepository, appID int, limit domain.Limit) error {
	limits := lor.Filter(func(l store.LimitOverride) bool {
		return l.AppID == appID && l.Limit.Key == limit.Key
	})

	switch len(limits) {
	case 0:
		lor.Create(store.LimitOverride{AppID: appID, Limit: limit})
	case 1:
		lor.Update(store.LimitOverride{ID: limits[0].ID, AppID: appID, Limit: limit})
	default:
		// This should never happen
		return errors.New("Duplicated limit override")
	}

	return nil
}

func (c *AppController) OptOutPublic(appID int) {
	app, err := c.ar.Get(appID)
	if err != nil {
		return
	}

	currentSub, err := c.sr.Get(app.SubscriptionID)
	if err != nil {
		return
	}

	targetSub, ok := findSubscription(c.sr, currentSub.UserID, false)
	if !ok {
		// This should never happen
		return
	}

	app.SubscriptionID = targetSub.ID
	c.ar.Update(app)
}

func (c *AppController) GetLimits(appID int) ([]domain.Limit, error) {
	app, err := c.ar.Get(appID)
	if err != nil {
		return []domain.Limit{}, err
	}

	sub, err := c.sr.Get(app.SubscriptionID)
	if err != nil {
		return []domain.Limit{}, err
	}

	plan, _ := c.pr.Get(sub.PlanID)

	overrides := func() []domain.Limit {
		limits := c.lor.Filter(func(l store.LimitOverride) bool {
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
