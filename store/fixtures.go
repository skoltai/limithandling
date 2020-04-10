package store

import "github.com/skoltai/limithandling/domain"

var Plans = planCollection{
	Plan{
		ID: 1,
		Plan: domain.Plan{
			Name:  "free",
			Price: 0,
			Limits: []domain.Limit{
				{Key: "concurrency", Value: 1},
				{Key: "buildtime", Value: 10},
				{Key: "builds", Value: 200},
				{Key: "teammembers", Value: 2},
			},
		},
	},
	Plan{
		ID: 2,
		Plan: domain.Plan{
			Name:  "Developer",
			Price: 40,
			Limits: []domain.Limit{
				{Key: "concurrency", Value: 2},
				{Key: "buildtime", Value: 45},
				{Key: "builds", Value: 0},
				{Key: "teammembers", Value: 0},
			},
		},
	},
	Plan{
		ID: 3,
		Plan: domain.Plan{
			Name:  "Organization",
			Price: 100,
			Limits: []domain.Limit{
				{Key: "concurrency", Value: 4},
				{Key: "buildtime", Value: 90},
				{Key: "builds", Value: 0},
				{Key: "teammembers", Value: 0},
			},
		},
	},
	Plan{
		ID: 4,
		Plan: domain.Plan{
			Name:  "Public",
			Price: 0,
			Limits: []domain.Limit{
				{Key: "concurrency", Value: 2},
				{Key: "buildtime", Value: 45},
				{Key: "builds", Value: 0},
				{Key: "teammembers", Value: 0},
			},
		},
	},
}

func NewTestStore() Store {
	return &MemoryStore{
		Users:          newUserCollection(),
		Plans:          &Plans,
		Subscriptions:  newSubscriptionCollection(),
		Apps:           newAppCollection(),
		LimitOverrides: newLimitOverrideCollection(),
	}
}
