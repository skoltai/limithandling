package controller

import (
	"testing"

	"github.com/skoltai/limithandling/domain"
	"github.com/skoltai/limithandling/store"
	"github.com/stretchr/testify/assert"
)

func TestCreateApp(t *testing.T) {
	s := store.NewTestStore()
	ac := NewAccountController(s)
	ac.Create(domain.User{Username: "testuser", Email: "testuser@example.com"}, 1)

	c := AppController{store: s}
	apps := []domain.App{
		domain.App{Name: "private-1", Public: false},
		domain.App{Name: "public-1", Public: true},
		domain.App{Name: "private-2", Public: false},
		domain.App{Name: "public-2", Public: true},
	}
	for _, a := range apps {
		c.Create(1, a)
	}

	sub, _ := s.GetSubscription(2)
	want := store.Subscription{
		ID: 2,
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