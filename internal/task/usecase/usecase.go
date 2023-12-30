package usecase

import (
	"context"
	"english_bot_admin/internal/module"
	"english_bot_admin/internal/task"
	"github.com/google/uuid"
)

type TaskUsecase struct {
	moduleUC module.Usecase
	taskRepo task.Repository
}

func NewTaskUsecase(moduleUc module.Usecase, taskRepo task.Repository) task.Usecase {
	return &TaskUsecase{
		moduleUC: moduleUc,
		taskRepo: taskRepo,
	}
}

func (u *TaskUsecase) AddToModule(params *task.ToModule) error {
	var (
		module_ *module.Module
		err     error
	)
	module_, err = u.moduleUC.GetModuleByID(params.ModuleID)
	if err != nil {

	}
	if module_.Task == nil {
		module_.Task = &[]uuid.UUID{}
	}

	*module_.Task = append(*module_.Task, params.TaskID)

	return nil
}

func (u *TaskUsecase) GetTaskById(context_ context.Context, uuid_ uuid.UUID) (*task.Task, error) {
	id, err := u.taskRepo.GetTaskByID(context_, uuid_)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (u *TaskUsecase) CreateTask(ctx context.Context, task *task.Task) (uuid.UUID, error) {
	uuid_, err := u.taskRepo.InsertTask(ctx, task)
	if err != nil {
		return uuid.UUID{}, err
	}
	return uuid_, nil
}
