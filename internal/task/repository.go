package task

import (
	"context"
	"english_bot_admin/internal/models"
	"github.com/google/uuid"
)

type Repository interface {
	GetTaskByID(context context.Context, uuid uuid.UUID) (*models.Task, error)
	InsertTask(ctx context.Context, task *models.Task) (uuid.UUID, error)
	GetTasks(ctx context.Context) ([]models.Task, error)
	UpdateTaskInfoByUUID(ctx context.Context, task *models.Task) error
	GetTasksByLvl(ctx context.Context, lvl string) ([]models.Task, error)
}
