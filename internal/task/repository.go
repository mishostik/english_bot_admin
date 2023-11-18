package task

import (
	model "english_bot_admin/internal/task/models"
)

type Repository interface {
	GetTasks() ([]model.Task, error)
	GetTaskByID(taskID int) (*model.Task, error)
	NewTask(task *model.Task) error
	UpdateTask(taskID int, task *model.Task) error
	DeleteTask(taskID int) error
}
