package store

import (
	"testing"

	"github.com/skoltai/limithandling/domain"
	"github.com/stretchr/testify/assert"
)

func TestAppCollection(t *testing.T) {
	c := NewAppCollection()
	app := App{
		OwnerID:        1,
		SubscriptionID: 2,
		App:            domain.App{Name: "testapp"},
	}
	app.ID = c.Create(app)

	got, err := c.Get(app.ID)
	assert.NoError(t, err)
	assert.Equal(t, app, got)
}
