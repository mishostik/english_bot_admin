package http

import (
	"encoding/json"
	model "english_bot_admin/internal/task/model"
	tr "english_bot_admin/internal/task/repository"
	uc "english_bot_admin/internal/task/usecase"

	"go.mongodb.org/mongo-driver/mongo"

	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	useCase  *uc.TaskUsecase
	taskRepo *tr.MongoTaskRepository
}

func NewTaskHandler(taskCollection *mongo.Collection, typeCollection *mongo.Collection) *TaskHandler {
	return &TaskHandler{
		useCase:  uc.NewTaskUsecase(),
		taskRepo: tr.NewMongoTaskRepository(taskCollection, typeCollection),
	}
}

func (h *TaskHandler) GetTasks(ctx *fiber.Ctx) error {
	context_ := ctx.Context()
	tasks, err := h.taskRepo.GetTasks(context_)
	if err != nil {
		return nil
	}
	jsonTasks, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return ctx.Render("../templates/tasks.html", jsonTasks)
}

func (h *TaskHandler) CreateTask(ctx *fiber.Ctx) error {
	context_ := ctx.Context()
	taskType := ctx.FormValue("type")
	level := ctx.FormValue("level")
	question := ctx.FormValue("question")
	answer := ctx.FormValue("answer")
	log.Println(taskType, level, question, answer)

	taskTypeInt, err := strconv.Atoi(taskType)

	newTask := &model.Task{
		TypeID:   uint8(taskTypeInt),
		Level:    level,
		Question: question,
		Answer:   answer,
	}
	err = h.taskRepo.NewTask(context_, newTask)
	if err != nil {
		return err
	}
	if err != nil {
		return nil
	}
	return ctx.SendString("Task created successfully!")
}

func (h *TaskHandler) UpdateTask(ctx *fiber.Ctx) error {
	context_ := ctx.Context()
	task := &model.Task{
		// updaload fields from form
	}
	err := h.taskRepo.UpdateTask(context_, 1, task)
	if err != nil {
		log.Println(err.Error())
	}
	return nil
}

func (h *TaskHandler) GetTaskByID(ctx *fiber.Ctx) error {
	context_ := ctx.Context()
	taskID := ctx.Params("id")

	id, err := strconv.Atoi(taskID)
	if err != nil {
		return err
	}

	task, err := h.taskRepo.GetTaskByID(context_, id)
	if err != nil {
		log.Println(err.Error())
	}

	jsonTask, err := json.Marshal(task)
	if err != nil {
		return err
	}

	return ctx.Render("../templates/task.html", jsonTask)
}

func (h *TaskHandler) DeleteTask(ctx *fiber.Ctx) error {
	var (
		err error
		id  int
	)
	context_ := ctx.Context()
	taskID := ctx.Params("id")

	id, err = strconv.Atoi(taskID)
	if err != nil {
		return err
	}

	err = h.taskRepo.DeleteTask(context_, id)
	if err != nil {
		log.Println(err.Error())
	}
	return ctx.SendString("Task deleted successfully")
}
