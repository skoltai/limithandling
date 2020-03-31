package store

import "github.com/skoltai/limithandling/domain"

var Plans = PlanCollection{
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
	Plan{
		ID: 2,
		Plan: domain.Plan{
			Name:  "Developer",
			Price: 40,
			Limits: []domain.Limit{
				domain.Limit{Key: "concurrency", Value: 2},
				domain.Limit{Key: "buildtime", Value: 45},
				domain.Limit{Key: "builds", Value: 0},
				domain.Limit{Key: "teammembers", Value: 0},
			},
		},
	},
	Plan{
		ID: 3,
		Plan: domain.Plan{
			Name:  "Organization",
			Price: 100,
			Limits: []domain.Limit{
				domain.Limit{Key: "concurrency", Value: 4},
				domain.Limit{Key: "buildtime", Value: 90},
				domain.Limit{Key: "builds", Value: 0},
				domain.Limit{Key: "teammembers", Value: 0},
			},
		},
	},
	Plan{
		ID: 4,
		Plan: domain.Plan{
			Name:  "Public",
			Price: 0,
			Limits: []domain.Limit{
				domain.Limit{Key: "concurrency", Value: 2},
				domain.Limit{Key: "buildtime", Value: 45},
				domain.Limit{Key: "builds", Value: 0},
				domain.Limit{Key: "teammembers", Value: 0},
			},
		},
	},
}
