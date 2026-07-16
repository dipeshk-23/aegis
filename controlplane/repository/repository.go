package repository

import (
	"context"

	governancev1alpha1 "github.com/dipeshk-23/aegis/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
)

type Repository interface {
	GetAgent(ctx context.Context, req ctrl.Request) (*governancev1alpha1.Agent, error)
}
