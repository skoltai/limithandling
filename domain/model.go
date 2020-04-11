// Package domain describes the (simplified) models used in the system and their relations to each other
package domain

// User represents a user in the system.
// It can have many Subscriptions.
type User struct {
	Username      string
	Email         string
	Subscriptions []Subscription
}

// Subscription represents a subscription in the system.
// It belongs to a User and a Plan, and can have many Apps
type Subscription struct {
	User   User
	Plan   Plan
	Public bool
	Apps   []App
}

// Plan represents an offering to which a User can subscribe.
// It has many Limits, which are the default limits for an App belonging to a Subscription with that Plan.
type Plan struct {
	Name   string
	Price  int
	Limits []Limit
}

// Limit represents an abstract resource limit in the system.
type Limit struct {
	Key   string
	Value int
}

// App represents a registered app in the system.
// It belongs to the owner User and to a Subscription which defines the default Limits which can be overridden by many LimitOverrides.
type App struct {
	Name          string
	Owner         User
	Public        bool
	Subscription  Subscription
	LimitOverride []Limit
	Builds        []Build
}

// Build represents a build job for an app and currently serves as a placeholder to illustrate the data model
type Build struct{}
