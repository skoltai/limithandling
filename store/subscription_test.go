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

	id := c.Create(s)
	s.ID = id

	got, err := c.Get(id)
	assert.NoError(t, err)

	assert.Equal(t, s, got)

	_, err = c.Get(2)
	assert.Error(t, err)
}
