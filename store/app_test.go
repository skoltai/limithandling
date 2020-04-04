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
	id := c.Create(app)
	app.ID = id

	got, err := c.Get(id)
	assert.NoError(t, err)
	assert.Equal(t, app, got)
}
