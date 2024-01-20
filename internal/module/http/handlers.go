package http

import (
	"bytes"
	"english_bot_admin/internal/httpServer/cconstants"
	"english_bot_admin/internal/models"
	"english_bot_admin/internal/module"
	"english_bot_admin/internal/task"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"html/template"
	"log"
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

func renderModules(ctx *fiber.Ctx, modules []models.Module) {
	tmpl, err := template.ParseFiles("templates/modules.html")
	if err != nil {
		err = ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		if err != nil {
			return
		}
		return
	}

	data := struct {
		Modules []models.Module
	}{
		Modules: modules,
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, data); err != nil {
		err = ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		if err != nil {
			return
		}
		return
	}
	ctx.Set("Content-Type", "text/html")
	err = ctx.Status(fiber.StatusOK).Send(buf.Bytes())
	if err != nil {
		return
	}
}

func (h *ModuleHandler) GetAllModules(ctx *fiber.Ctx) error {
	context_ := ctx.Context()
	modules, err := h.UC.GetModules(context_)
	if err != nil {
		err = ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		if err != nil {
			return err
		}
		return err
	}
	renderModules(ctx, modules)

	return nil
}

func (h *ModuleHandler) CreateModule(ctx *fiber.Ctx) error {
	var (
		context_ = ctx.Context()
		err      error
		params   models.NewModuleParams

		errorMessage string = cconstants.SuccessModuleAdd
	)

	if err = ctx.BodyParser(&params); err != nil {
		errorMessage = fmt.Sprintf("error parsing params: %v", err.Error())
	}

	if params.Title == "" {
		return fiber.NewError(fiber.StatusBadRequest, cconstants.TitleRequired)
	}
	if params.Level == "" {
		return fiber.NewError(fiber.StatusBadRequest, cconstants.LevelRequired)
	}

	err = h.UC.GenerateModule(context_, &params)
	if err != nil {
		errorMessage = err.Error()
	}

	data := fiber.Map{
		"Message": errorMessage,
	}

	return ctx.Render("templates/message.html", data)
}

func (h *ModuleHandler) GetNewModuleForm(ctx *fiber.Ctx) error {
	return ctx.Render("templates/create_module.html", fiber.Map{})
}

func renderTasks(ctx *fiber.Ctx, tasks []models.ByModule) {
	tmpl, err := template.ParseFiles("templates/tasks_by_lvl.html")
	if err != nil {
		err = ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		if err != nil {
			return
		}
		return
	}

	data := struct {
		Tasks []models.ByModule
	}{
		Tasks: tasks,
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, data); err != nil {
		err = ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		if err != nil {
			return
		}
		return
	}
	ctx.Set("Content-Type", "text/html")
	err = ctx.Status(fiber.StatusOK).Send(buf.Bytes())
	if err != nil {
		return
	}
}

func (h *ModuleHandler) GetTasksByLvl(ctx *fiber.Ctx) error {
	var (
		context_ = ctx.Context()
		params   = &models.ByLvl{}
		err      error
	)
	params.Level = ctx.Query("level")
	moduleIdStr := ctx.Query("module_id")
	moduleID, err := uuid.Parse(moduleIdStr)
	if err != nil {
		return fmt.Errorf("error getting params (module id): %v", err)
	}

	params.ModuleID = moduleID

	tasksByLvl, err := h.taskUC.GetTasksByLvl(context_, params)
	if err != nil {
		log.Println(err)
		return err
	}

	renderTasks(ctx, tasksByLvl)

	return nil
}

func (h *ModuleHandler) AddTasksByLvl(ctx *fiber.Ctx) error {
	var (
		params       models.TaskToModule
		context_            = ctx.Context()
		errorMessage string = "Task successfully added to module"
	)

	taskIDStr := ctx.Query("task_id")
	taskID, err := uuid.Parse(taskIDStr)
	if err != nil {
		//errorMessage = err.Error()
		err = ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		if err != nil {
			return err
		}
	}

	// todo: generate task
	// 1 - find by id
	receivedTask, err := h.taskUC.GetTaskById(context_, taskID)
	if err != nil {
		err = ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		if err != nil {
			return err
		}
	}

	if receivedTask == nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Task by uuid is not received"))
		//if err != nil {
		//	return err
		//}
	}

	// 2 - get the struct
	temp := &models.Task{
		TaskID:   taskID,
		TypeID:   receivedTask.TypeID,
		Level:    receivedTask.Level,
		Question: receivedTask.Question,
		Answer:   receivedTask.Answer,
	}

	// 3 - params.Task = temp struct
	params.Task = temp

	moduleIdStr := ctx.Query("module_id")
	moduleID, err := uuid.Parse(moduleIdStr)
	if err != nil {
		err = ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		if err != nil {
			return err
		}
	}

	params.ModuleId = moduleID

	err = h.UC.AddTask(context_, &params)
	if err != nil {
		err = ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		if err != nil {
			return err
		}
	}

	data := fiber.Map{
		"Message": errorMessage,
	}

	return ctx.Render("templates/message.html", data)
}
