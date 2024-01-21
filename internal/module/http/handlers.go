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

// GetAllModules
// @Summary GetAllModules
// @Description Get all modules
// @ID get all modules
// @Produce json
// @Success 200 {object} models.ModulesResponseModel
// @Failure 500 {object} models.ModulesResponseModel
// @Router /module/all [get]
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

// CreateModule
// @Summary CreateModule
// @Description Create module
// @ID create-module
// @Produce json
// @Param request body models.NewModuleParams true
// @Success 200 {object} models.ResponseModel
// @Failure 500 {object} models.ResponseModel
// @Router /module/new [post]
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

// GetTasksByLvl
// @Summary GetTasksByLvl
// @Description Get tasks by level
// @ID get-task-by-level
// @Produce json
// @Param request body models.ByLvl true
// @Success 200 {object} models.TasksByLvlResponseModel
// @Failure 500 {object} models.TasksByLvlResponseModel
// @Router /module/task/by_lvl [post]
func (h *ModuleHandler) GetTasksByLvl(ctx *fiber.Ctx) error {
	var (
		context_ = ctx.Context()
		params   models.ByLvl
		response *models.TasksByLvlResponseModel = &models.TasksByLvlResponseModel{}
		err      error
	)

	if err = ctx.BodyParser(&params); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	tasksByLvl, err := h.taskUC.GetTasksByLvl(context_, &params)
	if err != nil {
		response.Error = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Success = true
	response.Data = tasksByLvl
	return ctx.Status(fiber.StatusInternalServerError).JSON(response)
}

// AddTasksByLvl
// @Summary AddTasksByLvl
// @Description Add tasks to module by level
// @ID add-task-by-lvl
// @Produce json
// @Param request body models.AddTaskByLvlParams true
// @Success 200 {object} models.ResponseModel
// @Failure 500 {object} models.ResponseModel
// @Router /module/task/add [post]
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
