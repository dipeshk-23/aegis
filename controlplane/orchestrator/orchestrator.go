package orchestrator

import (
	"context"
	"fmt"

	"github.com/dipeshk-23/aegis/controlplane/repository"
	ctrl "sigs.k8s.io/controller-runtime"
)

type Orchestrator struct {
	repository repository.Repository
}

func New(r repository.Repository) *Orchestrator {
	return &Orchestrator{
		repository: r,
	}
}
func (o *Orchestrator) Handle(ctx context.Context, req ctrl.Request) error {
	agent, err := o.repository.GetAgent(ctx, req)
	if err != nil {
		return err
	}
	fmt.Println(agent.Spec.DisplayName)
	return nil
}

func (o *Orchestrator) Reconcile() error {
	//fmt.Println("Aegis Orchestrator Invoked")
	return nil
}
