package domain

var Plans = []Plan{
	Plan{
		Name:  "free",
		Price: 0,
		Limits: []Limit{
			Limit{Key: "concurrency", Value: 1},
			Limit{Key: "buildtime", Value: 10},
			Limit{Key: "builds", Value: 200},
			Limit{Key: "teammembers", Value: 2},
		},
	},
	Plan{
		Name:  "Developer",
		Price: 40,
		Limits: []Limit{
			Limit{Key: "concurrency", Value: 2},
			Limit{Key: "buildtime", Value: 45},
			Limit{Key: "builds", Value: 0},
			Limit{Key: "teammembers", Value: 0},
		},
	},
	Plan{
		Name:  "Organization",
		Price: 100,
		Limits: []Limit{
			Limit{Key: "concurrency", Value: 4},
			Limit{Key: "buildtime", Value: 90},
			Limit{Key: "builds", Value: 0},
			Limit{Key: "teammembers", Value: 0},
		},
	},
	Plan{
		Name:  "Public",
		Price: 0,
		Limits: []Limit{
			Limit{Key: "concurrency", Value: 2},
			Limit{Key: "buildtime", Value: 45},
			Limit{Key: "builds", Value: 0},
			Limit{Key: "teammembers", Value: 0},
		},
	},
}
