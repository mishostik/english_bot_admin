package task

import (
	"context"
	"github.com/google/uuid"
)

type Usecase interface {
	AddToModule(params *ToModule) error
	GetTaskById(context_ context.Context, uuid_ uuid.UUID) (*Task, error)
	CreateTask(ctx context.Context, task *Task) (uuid.UUID, error)
}
