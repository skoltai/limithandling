package controller

import (
	"testing"

	"github.com/skoltai/limithandling/domain"
	"github.com/skoltai/limithandling/store"
	"github.com/stretchr/testify/assert"
)

func TestAccountCreate(t *testing.T) {
	s := store.NewTestStore()
	sr := store.NewSimpleSubscriptionRepository(s)
	c := NewAccountController(store.NewSimpleUserRepository(s), sr)
	user := domain.User{Username: "testuser", Email: "testuser@example.com"}
	planID := 1

	c.Create(user, planID)
	sub, _ := sr.Get(1)

	assert.Equal(t, 1, sub.UserID)
	assert.Equal(t, planID, sub.PlanID)
}
