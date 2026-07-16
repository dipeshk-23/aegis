package kubernetes

import (
	//"context"

	"context"

	governancev1alpha1 "github.com/dipeshk-23/aegis/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Repository struct {
	client.Client
}

func New(c client.Client) *Repository {
	return &Repository{
		Client: c,
	}
}
func (r *Repository) GetAgent(ctx context.Context, req ctrl.Request) (*governancev1alpha1.Agent, error) {
	agent := &governancev1alpha1.Agent{}

	err := r.Get(
		ctx,
		req.NamespacedName,
		agent,
	)
	if err != nil {
		return nil, err
	}

	return agent, nil
}
