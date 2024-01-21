package module

import (
	"context"
	"english_bot_admin/internal/models"
	"github.com/google/uuid"
)

type Usecase interface {
	GenerateModule(ctx context.Context, newModule *models.NewModuleParams) error
	ChangeModule(newTasksNum []uuid.UUID) error
	GetModules(ctx context.Context) ([]models.Module, error)
	GetModuleByID(moduleID uuid.UUID) (*models.Module, error)
	AddTask(ctx context.Context, params *models.TaskToModule) error
}
