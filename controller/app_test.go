package controller

import (
	"testing"

	"github.com/skoltai/limithandling/domain"
	"github.com/skoltai/limithandling/store"
	"github.com/stretchr/testify/assert"
)

func TestCreateApp(t *testing.T) {
	s := store.NewTestStore()

	sr := store.NewSimpleSubscriptionRepository(s)
	c := NewAppController(sr, store.NewSimpleAppRepository(s), store.NewSimpleLimitOverrideRepository(s), store.NewSimplePlanRepository(s))

	err := c.Create(1, domain.App{})
	assert.Error(t, err)
	
	ac := NewAccountController(store.NewSimpleUserRepository(s), store.NewSimpleSubscriptionRepository(s))
	ac.Create(domain.User{Username: "testuser", Email: "testuser@example.com"}, 1)

	apps := []domain.App{
		{Name: "private-1", Public: false},
		{Name: "public-1", Public: true},
		{Name: "private-2", Public: false},
		{Name: "public-2", Public: true},
	}
	for _, a := range apps {
		c.Create(1, a)
	}

	sub, _ := sr.Get(2)
	want := store.Subscription{
		ID:     2,
		UserID: 1,
		PlanID: 4, // bad magic number
		Subscription: domain.Subscription{
			Public: true,
		},
	}
	assert.Equal(t, want, sub)

	got := func() []domain.App {
		apps := make([]domain.App, 0)
		for _, a := range s.GetApps() {
			apps = append(apps, a.App)
		}
		return apps
	}()

	assert.Greater(t, len(got), 0)
	assert.ElementsMatch(t, apps, got)
}

func TestSetCustomLimits(t *testing.T) {
	s := store.NewTestStore()
	ar, lor := store.NewSimpleAppRepository(s), store.NewSimpleLimitOverrideRepository(s)
	c := NewAppController(store.NewSimpleSubscriptionRepository(s), ar, lor, store.NewSimplePlanRepository(s))

	app := store.App{
		OwnerID:        1,
		SubscriptionID: 2,
		App:            domain.App{Name: "testapp"},
	}
	app.ID = ar.Create(app)

	limits := []domain.Limit{
		{Key: "concurrency", Value: 1},
		{Key: "buildtime", Value: 10},
		{Key: "builds", Value: 200},
		{Key: "concurrency", Value: 2},
		{Key: "buildtime", Value: 45},
		{Key: "builds", Value: 0},
		{Key: "teammembers", Value: 0},
	}

	getLimitOverrides := func() []domain.Limit {
		l := make([]domain.Limit, 0)
		for _, lo := range lor.Filter(func(l store.LimitOverride) bool {
			return l.AppID == 1
		}) {
			l = append(l, lo.Limit)
		}
		return l
	}

	c.SetCustomLimits(app.ID, limits[0:3])
	assert.ElementsMatch(t, limits[0:3], getLimitOverrides())

	c.SetCustomLimits(app.ID, limits[3:7])
	assert.ElementsMatch(t, limits[3:7], getLimitOverrides())

	c.SetCustomLimits(app.ID, []domain.Limit{limits[2], limits[5], limits[2], limits[5], limits[5]})
	assert.ElementsMatch(t, limits[3:7], getLimitOverrides())
}

func TestOptOutPublic(t *testing.T) {
	s := store.NewTestStore()
	sr := store.NewSimpleSubscriptionRepository(s)
	ac := NewAccountController(store.NewSimpleUserRepository(s), sr)
	ac.Create(domain.User{Username: "testuser", Email: "testuser@example.com"}, 1)

	c := NewAppController(sr, store.NewSimpleAppRepository(s), store.NewSimpleLimitOverrideRepository(s), store.NewSimplePlanRepository(s))
	apps := []domain.App{
		{Name: "private-1", Public: false},
		{Name: "public-1", Public: true},
		{Name: "private-2", Public: false},
		{Name: "public-2", Public: true},
	}
	for _, a := range apps {
		c.Create(1, a)
	}

	appIDsInPrivateSubscription := func() []int {
		res := make([]int, 0)
		sub, ok := findSubscription(sr, 1, false)
		if !ok {
			return res
		}
		for _, app := range s.GetApps() {
			if app.SubscriptionID == sub.ID {
				res = append(res, app.ID)
			}
		}
		return res
	}

	assert.ElementsMatch(t, []int{1, 3}, appIDsInPrivateSubscription())

	c.OptOutPublic(1)
	c.OptOutPublic(2)
	c.OptOutPublic(3)

	assert.ElementsMatch(t, []int{1, 2, 3}, appIDsInPrivateSubscription())
}

func TestGetLimits(t *testing.T) {
	s := store.NewTestStore()
	sr := store.NewSimpleSubscriptionRepository(s)
	ac := NewAccountController(store.NewSimpleUserRepository(s), sr)
	ac.Create(domain.User{Username: "testuser", Email: "testuser@example.com"}, 1)

	c := NewAppController(sr, store.NewSimpleAppRepository(s), store.NewSimpleLimitOverrideRepository(s), store.NewSimplePlanRepository(s))
	apps := []domain.App{
		{Name: "private-1", Public: false},
		{Name: "public-1", Public: true},
		{Name: "private-2", Public: false},
		{Name: "public-2", Public: true},
	}
	for _, a := range apps {
		c.Create(1, a)
	}

	want := []domain.Limit{
		{Key: "concurrency", Value: 2},
		{Key: "buildtime", Value: 10},
		{Key: "builds", Value: 200},
		{Key: "teammembers", Value: 0},
	}

	c.SetCustomLimits(1, []domain.Limit{want[0], want[3]})

	got, _ := c.GetLimits(1)
	assert.ElementsMatch(t, want, got)
}
