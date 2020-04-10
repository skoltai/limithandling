package store

import (
	"testing"

	"github.com/skoltai/limithandling/domain"
	"github.com/stretchr/testify/assert"
)

func TestSubscriptionCollection(t *testing.T) {
	c := newSubscriptionCollection()

	s := Subscription{
		UserID:       1,
		PlanID:       1,
		Subscription: domain.Subscription{Public: true},
	}

	s.ID = c.create(s)

	got, err := c.get(s.ID)
	assert.NoError(t, err)

	assert.Equal(t, s, got)

	_, err = c.get(2)
	assert.Error(t, err)
}
