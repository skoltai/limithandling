package store

type MemoryStore struct {
	Users          *userCollection
	Subscriptions  *subscriptionCollection
	Plans          *planCollection
	Apps           *appCollection
	LimitOverrides *limitOverrideCollection
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		Users:          newUserCollection(),
		Subscriptions:  newSubscriptionCollection(),
		Apps:           newAppCollection(),
		LimitOverrides: newLimitOverrideCollection(),
	}
}

func (s *MemoryStore) GetApps() map[int]App {
	return s.Apps.items
}
