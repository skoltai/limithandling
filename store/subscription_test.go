package store

import (
	"testing"

	"github.com/skoltai/limithandling/domain"
	"github.com/stretchr/testify/assert"
)

func TestSubscriptionCollection(t *testing.T) {
	c := NewSubscriptionCollection()

	s := Subscription{
		UserID:       1,
		PlanID:       1,
		Subscription: domain.Subscription{Public: true},
	}

	s.ID = c.Create(s)

	got, err := c.Get(s.ID)
	assert.NoError(t, err)

	assert.Equal(t, s, got)

	_, err = c.Get(2)
	assert.Error(t, err)
}
