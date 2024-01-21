package module

import (
	"context"
	"english_bot_admin/internal/models"
)

type Repository interface {
	NewModule(ctx context.Context, module *models.Module) error
	SelectModules(ctx context.Context) ([]models.Module, error)
	InsertTask(ctx context.Context, params *models.TaskToModule) error
}
