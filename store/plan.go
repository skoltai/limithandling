package store

import (
	"errors"

	"github.com/skoltai/limithandling/domain"
)

type Plan struct {
	domain.Plan
	ID int
}

type planCollection []Plan

func (p planCollection) get(id int) (Plan, error) {
	for _, plan := range p {
		if plan.ID == id {
			return plan, nil
		}
	}

	return Plan{}, errors.New("Plan not found")
}
