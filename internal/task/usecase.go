package task

import (
	"context"
	"english_bot_admin/internal/models"
	"github.com/google/uuid"
)

type Usecase interface {
	GetTasks(context_ context.Context) ([]models.Task, error)
	GetTaskById(context_ context.Context, uuid_ uuid.UUID) (*models.Task, error)
	CreateTask(ctx context.Context, task *models.Task) (uuid.UUID, error)
	GetTasksByLvl(ctx context.Context, params *models.ByLvl) ([]models.ByModule, error)
	UpdateTaskInfoByUUID(ctx context.Context, task *models.Task) error
}
