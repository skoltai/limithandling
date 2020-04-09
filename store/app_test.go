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

func TestAppUpdate(t *testing.T) {
	c := NewAppCollection()
	assert.False(t, c.Update(App{ID: 0}))
	assert.False(t, c.Update(App{ID: 1}))

	app := App{
		OwnerID:        1,
		SubscriptionID: 1,
		App:            domain.App{Name: "test1"},
	}
	app.ID = c.Create(app)

	got, _ := c.Get(app.ID)
	assert.Equal(t, app, got)

	want := App{
		ID:             app.ID,
		OwnerID:        1,
		SubscriptionID: 2,
		App:            domain.App{Name: "test1"},
	}

	ok := c.Update(want)
	assert.True(t, ok)
	got, _ = c.Get(want.ID)
	assert.Equal(t, want, got)
}
