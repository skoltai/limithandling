package store

import (
	"testing"

	"github.com/skoltai/limithandling/domain"
	"github.com/stretchr/testify/assert"
)

func TestAppCollection(t *testing.T) {
	c := newAppCollection()
	app := App{
		OwnerID:        1,
		SubscriptionID: 2,
		App:            domain.App{Name: "testapp"},
	}
	app.ID = c.create(app)

	got, err := c.get(app.ID)
	assert.NoError(t, err)
	assert.Equal(t, app, got)

	_, err = c.get(0)
	assert.Error(t, err)
}

func TestAppUpdate(t *testing.T) {
	c := newAppCollection()
	assert.False(t, c.update(App{ID: 0}))
	assert.False(t, c.update(App{ID: 1}))

	app := App{
		OwnerID:        1,
		SubscriptionID: 1,
		App:            domain.App{Name: "test1"},
	}
	app.ID = c.create(app)

	got, _ := c.get(app.ID)
	assert.Equal(t, app, got)

	want := App{
		ID:             app.ID,
		OwnerID:        1,
		SubscriptionID: 2,
		App:            domain.App{Name: "test1"},
	}

	ok := c.update(want)
	assert.True(t, ok)
	got, _ = c.get(want.ID)
	assert.Equal(t, want, got)
}
