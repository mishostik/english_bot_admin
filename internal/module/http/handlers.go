package http

import (
	"english_bot_admin/internal/models"
	"english_bot_admin/internal/module"
	"english_bot_admin/internal/task"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type ModuleHandler struct {
	UC     module.Usecase
	taskUC task.Usecase
}

func NewModuleHandler(useCase module.Usecase, taskUC task.Usecase) *ModuleHandler {
	return &ModuleHandler{
		UC:     useCase,
		taskUC: taskUC,
	}
}

func (h *ModuleHandler) GetAllModules(ctx *fiber.Ctx) error {
	var (
		context_                              = ctx.Context()
		response *models.ModulesResponseModel = &models.ModulesResponseModel{}
	)
	modules, err := h.UC.GetModules(context_)
	if err != nil {
		response.Error = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Success = true
	response.Data = modules
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (h *ModuleHandler) CreateModule(ctx *fiber.Ctx) error {
	var (
		context_ = ctx.Context()
		err      error
		params   models.NewModuleParams

		response *models.ResponseModel = &models.ResponseModel{}
	)

	if err = ctx.BodyParser(&params); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	err = h.UC.GenerateModule(context_, &params)
	if err != nil {
		response.Error = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Success = true
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (h *ModuleHandler) GetTasksByLvl(ctx *fiber.Ctx) error {
	var (
		context_ = ctx.Context()
		params   *models.ByLvl
		response *models.TasksByLvlResponseModel = &models.TasksByLvlResponseModel{}
		err      error
	)

	if err = ctx.BodyParser(&params); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	tasksByLvl, err := h.taskUC.GetTasksByLvl(context_, params)
	if err != nil {
		response.Error = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Success = true
	response.Data = tasksByLvl
	return ctx.Status(fiber.StatusInternalServerError).JSON(response)
}

func (h *ModuleHandler) AddTasksByLvl(ctx *fiber.Ctx) error {
	var (
		params          models.AddTaskByLvlParams
		paramsToAddTask models.TaskToModule
		context_                              = ctx.Context()
		response        *models.ResponseModel = &models.ResponseModel{}

		err error
	)

	if err = ctx.BodyParser(&params); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	receivedTask, err := h.taskUC.GetTaskById(context_, params.TaskId)
	if err != nil {
		response.Error = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	if receivedTask == nil {
		response.Error = fmt.Sprintf("error getting the task")
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	temp := &models.Task{
		TaskID:   params.TaskId,
		TypeID:   receivedTask.TypeID,
		Level:    receivedTask.Level,
		Question: receivedTask.Question,
		Answer:   receivedTask.Answer,
	}

	paramsToAddTask.Task = temp
	paramsToAddTask.ModuleId = params.ModuleId

	err = h.UC.AddTask(context_, &paramsToAddTask)
	if err != nil {
		response.Error = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
