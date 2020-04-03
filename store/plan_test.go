package store

import (
	"testing"

	"github.com/skoltai/limithandling/domain"
	"github.com/stretchr/testify/assert"
)

func TestPlanCollection(t *testing.T) {
	plans := PlanCollection{
		Plan{
			ID: 1,
			Plan: domain.Plan{
				Name:  "free",
				Price: 0,
				Limits: []domain.Limit{
					domain.Limit{Key: "concurrency", Value: 1},
					domain.Limit{Key: "buildtime", Value: 10},
					domain.Limit{Key: "builds", Value: 200},
					domain.Limit{Key: "teammembers", Value: 2},
				},
			},
		},
		Plan{ID: 2, Plan: domain.Plan{Name: "test"}},
	}

	got, err := plans.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, plans[0], got)

	_, err = plans.Get(3)
	assert.Error(t, err)
}
