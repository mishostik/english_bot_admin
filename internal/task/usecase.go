package task

type UseCase interface {
	GetTasks()
	GetTaskByID()
	NewTask()
	UpdateTask()
}
