package task

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	//GetTasks() ([]Task, error)
	GetTaskByID(context context.Context, uuid uuid.UUID) (*Task, error)
	InsertTask(ctx context.Context, task *Task) (uuid.UUID, error)
	GetTasks(ctx context.Context) ([]Task, error)
	//GetTasks(context_ *fasthttp.RequestCtx) (interface{}, interface{})
	UpdateTaskInfoByUUID(ctx context.Context, task *Task) error
	//UpdateTask(taskID int, task *Task) error
	//DeleteTask(taskID int) error
}
