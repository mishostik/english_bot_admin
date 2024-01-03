package module

import (
	"context"
)

type Repository interface {
	NewModule(ctx context.Context, module *Module) error
	SelectModules(ctx context.Context) ([]Module, error)
	InsertTask(ctx context.Context, params TaskToModule) error
}
