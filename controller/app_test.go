package controller

import (
	"testing"

	"github.com/skoltai/limithandling/domain"
	"github.com/skoltai/limithandling/store"
	//"github.com/stretchr/testify/assert"
)

func TestCreateApp(t *testing.T) {
	store := store.NewTestStore()
	ac := NewAccountController(store)
	ac.Create(domain.User{Username: "testuser", Email: "testuser@example.com"}, 1)

	c := AppController{store: store}
	c.Create(1, domain.App{Name: "private-1", Public: false})
	c.Create(1, domain.App{Name: "public-1", Public: true})
	c.Create(1, domain.App{Name: "private-2", Public: false})
	c.Create(1, domain.App{Name: "public-2", Public: true})
}