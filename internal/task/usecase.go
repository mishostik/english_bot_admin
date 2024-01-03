package task

import (
	"context"
	"english_bot_admin/internal/module"
	"github.com/google/uuid"
)

type Usecase interface {
	GetTasks(context_ context.Context) ([]Task, error)
	GetTaskById(context_ context.Context, uuid_ uuid.UUID) (*Task, error)
	CreateTask(ctx context.Context, task *Task) (uuid.UUID, error)
	GetTasksByLvl(ctx context.Context, params module.Lvl) ([]ByModule, error)
	UpdateTaskInfoByUUID(ctx context.Context, task *Task) error
}
