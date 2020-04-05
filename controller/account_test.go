package controller

import (
	"testing"

	"github.com/skoltai/limithandling/domain"
	"github.com/skoltai/limithandling/store"
	"github.com/stretchr/testify/assert"
)

func TestAccountCreate(t *testing.T) {
	c := AccountController{store: store.NewTestStore()}
	user := domain.User{Username: "testuser", Email: "testuser@example.com"}
	planID := 1

	c.Create(user, planID)
	s, _ := c.store.GetSubscription(1)

	assert.Equal(t, 1, s.UserID)
	assert.Equal(t, planID, s.PlanID)
}
