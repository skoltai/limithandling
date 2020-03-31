package store

import (
	"errors"

	"github.com/skoltai/limithandling/domain"
)

type Plan struct {
	domain.Plan
	ID int
}

type PlanCollection []Plan

func (p PlanCollection) Get(id int) (Plan, error) {
	for _, plan := range p {
		if plan.ID == id {
			return plan, nil
		}
	}

	return Plan{}, errors.New("Plan not found")
}
