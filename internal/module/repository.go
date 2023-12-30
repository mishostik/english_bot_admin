package module

import (
	"context"
)

type Repository interface {
	NewModule(ctx context.Context, module *Module) error
	SelectModules(ctx context.Context) ([]Module, error)
}
