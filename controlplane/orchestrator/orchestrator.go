package orchestrator

import (
	"context"
	"fmt"

	"github.com/dipeshk-23/aegis/controlplane/governance"
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

	// integrate decision engine
	decisionEngine := governance.New()
	decision, err := decisionEngine.Evaluate(agent)
	if err != nil {
		return err
	}
	fmt.Println("Agent :", agent.Name)
	fmt.Println("Display:", agent.Spec.DisplayName)
	fmt.Println("Trust Score:", decision.TrustScore)
	fmt.Println("State:", decision.Status)
	return nil
}

func (o *Orchestrator) Reconcile() error {
	//fmt.Println("Aegis Orchestrator Invoked")
	return nil
}
