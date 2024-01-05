package usecase

import (
	"context"
	"english_bot_admin/internal/task"
	"fmt"
	"github.com/google/uuid"
)

type TaskUsecase struct {
	repo task.Repository
}

func NewTaskUsecase(repo task.Repository) task.Usecase {
	return &TaskUsecase{
		repo: repo,
	}
}

func (u *TaskUsecase) GetTasks(context_ context.Context) ([]task.Task, error) {
	tasks, err := u.repo.GetTasks(context_)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (u *TaskUsecase) GetTaskById(context_ context.Context, uuid_ uuid.UUID) (*task.Task, error) {
	task_, err := u.repo.GetTaskByID(context_, uuid_)
	if err != nil {
		return nil, err
	}
	return task_, nil
}

func (u *TaskUsecase) CreateTask(ctx context.Context, task *task.Task) (uuid.UUID, error) {
	uuid_, err := u.repo.InsertTask(ctx, task)
	if err != nil {
		return uuid.UUID{}, err
	}
	return uuid_, nil
}

func (u *TaskUsecase) GetTasksByLvl(ctx context.Context, params *task.ByLvl) ([]task.ByModule, error) {
	var (
		fullTasks []task.Task
		tasks     []task.ByModule
		err       error
	)

	fullTasks, err = u.repo.GetTasksByLvl(ctx, params.Level)
	if err != nil {
		fmt.Println("full tasks and err:", fullTasks, err)
		return []task.ByModule{}, err
	}

	if len(fullTasks) == 0 {
		return []task.ByModule{}, fmt.Errorf("task by level {%s} not found", params.Level)
	}

	for _, task_ := range fullTasks {
		temp := &task.ByModule{
			ModuleID: params.ModuleID,
			TaskID:   task_.TaskID,
			Question: task_.Question,
			TypeID:   task_.TypeID,
		}
		tasks = append(tasks, *temp)
	}
	return tasks, nil
}

func (u *TaskUsecase) UpdateTaskInfoByUUID(ctx context.Context, task *task.Task) error {
	err := u.repo.UpdateTaskInfoByUUID(ctx, task)
	if err != nil {
		return err
	}
	return nil
}
