package domain

type User struct {
	Username      string
	Email         string
	Subscriptions []Subscription
}

type Subscription struct {
	User   User
	Plan   Plan
	Public bool
	Apps   []App
}

type Plan struct {
	Name   string
	Price  int
	Limits []Limit
}

type Limit struct {
	Key   string
	Value int
}

type App struct {
	Name          string
	Owner         User
	Public        bool
	Subscription  Subscription
	LimitOverride []Limit
	Builds        []Build
}

type Build struct{}
