package module

import (
	"context"
	"github.com/google/uuid"
)

type Usecase interface {
	GenerateModule(ctx context.Context, newModule NewModuleParams) error
	ChangeModule(newTasksNum []uuid.UUID) error
	GetModules(ctx context.Context) ([]Module, error)
	GetModuleByID(moduleID uuid.UUID) (*Module, error)
	AddTask(ctx context.Context, params TaskToModule) error
}
