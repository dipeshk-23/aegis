package governance

import (
	governancev1alpha1 "github.com/dipeshk-23/aegis/api/v1alpha1"
)

type Engine struct {
}

type Decision struct {
	TrustScore int
	Status     string
}

func New() *Engine {
	return &Engine{}
}
func (e *Engine) Evaluate(agent *governancev1alpha1.Agent) (*Decision, error) {
	return &Decision{
		TrustScore: 100,
		Status:     "Trusted",
	}, nil
}
