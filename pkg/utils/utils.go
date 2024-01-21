package utils

import (
	taskR "english_bot_admin/internal/task/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
)

type TaskHandler struct {
	taskRepo *taskR.MongoTaskRepository
}

func (h *TaskHandler) UpdateTasksUUID(ctx *fiber.Ctx) error {
	log.Println("uuid updating...")
	context_ := ctx.Context()

	tasks, err := h.taskRepo.GetTasksWithoutUUID(context_)
	if err != nil {
		return err
	}
	log.Println("tasks without uuid received")
	for _, task := range tasks {
		newUUID := uuid.New()
		task.TaskID = newUUID

		err = h.taskRepo.UpdateTaskUUID(context_, &task)
		if err != nil {
			return err
		}
	}
	log.Println("...tasks updated")
	return ctx.SendString("UUID успешно добавлен к задачам без UUID")
}
