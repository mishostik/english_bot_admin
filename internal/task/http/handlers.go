package http

import (
	"english_bot_admin/internal/incorrect"
	"english_bot_admin/internal/models"
	"english_bot_admin/internal/task"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	_ "github.com/swaggo/swag"
)

type TaskHandler struct {
	uc          task.Usecase
	taskRepo    task.Repository
	incorrectUC incorrect.Usecase
}

func NewTaskHandler(taskUC task.Usecase, taskRepo task.Repository, incUC incorrect.Usecase) *TaskHandler {
	return &TaskHandler{
		uc:          taskUC,
		taskRepo:    taskRepo,
		incorrectUC: incUC,
	}
}

// GetTasks
// @Summary GetTasks
// @Description Get all tasks
// @ID get-all-tasks
// @Produce json
// @Success 200 {object} models.TasksResponseModel
// @Failure 500 {object} models.TasksResponseModel
// @Router /task/all [get]
func (h *TaskHandler) GetTasks(ctx *fiber.Ctx) error {
	var (
		context_                            = ctx.Context()
		response *models.TasksResponseModel = &models.TasksResponseModel{}
	)
	tasks, err := h.uc.GetTasks(context_)
	if err != nil {
		response.Error = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Success = true
	response.Data = tasks
	return ctx.Status(fiber.StatusOK).JSON(response)
}

// EditTask
// @Summary EditTask
// @Description Edit task
// @ID edit-task
// @Produce json
// @Param request body models.TaskWithAnswers true
// @Success 200 {object} models.ResponseModel
// @Failure 500 {object} models.ResponseModel
// @Router /task/edit [post]
func (h *TaskHandler) EditTask(ctx *fiber.Ctx) error {
	var (
		context_ = ctx.Context()
		params   models.TaskWithAnswers
		response *models.ResponseModel = &models.ResponseModel{}

		err error
	)

	if err = ctx.BodyParser(&params); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	editTask := &models.Task{
		TaskID:   params.TaskID,
		TypeID:   params.TypeID,
		Level:    params.Level,
		Question: params.Question,
		Answer:   params.Answer,
	}

	err = h.uc.UpdateTaskInfoByUUID(context_, editTask)
	if err != nil {
		response.Error = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	incAnswers := &models.IncorrectAnswers{
		A: params.A,
		B: params.B,
		C: params.C,
	}

	err = h.incorrectUC.UpdateForTask(context_, params.TaskID, incAnswers)
	if err != nil {
		response.Error = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Success = true
	return ctx.SendStatus(fiber.StatusOK)
}

// CreateTask
// @Summary CreateTask
// @Description Create task
// @ID create-task
// @Produce json
// @Param request body models.TaskWithAnswers true
// @Success 200 {object} models.ResponseModel
// @Failure 500 {object} models.ResponseModel
// @Router /task/new [post]
func (h *TaskHandler) CreateTask(ctx *fiber.Ctx) error {
	var (
		context_   = ctx.Context()
		params     models.TaskWithAnswers
		response   *models.ResponseModel = &models.ResponseModel{}
		internalId uuid.UUID
		err        error
	)

	if err = ctx.BodyParser(&params); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	newTask := &models.Task{
		TaskID:   uuid.New(),
		TypeID:   params.TypeID,
		Level:    params.Level,
		Question: params.Question,
		Answer:   params.Answer,
	}

	internalId, err = h.uc.CreateTask(context_, newTask)
	if err != nil {
		response.Error = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	err = h.incorrectUC.AddForTask(context_, internalId, params.A, params.B, params.C)
	if err != nil {
		response.Error = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
