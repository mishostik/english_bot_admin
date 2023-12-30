package http

import (
	"bytes"
	"english_bot_admin/internal/module"
	"github.com/gofiber/fiber/v2"
	"html/template"
)

type ModuleHandler struct {
	UC module.Usecase
}

func NewModuleHandler(useCase module.Usecase) *ModuleHandler {
	return &ModuleHandler{
		UC: useCase,
	}
}

func renderModules(ctx *fiber.Ctx, modules []module.Module) {
	tmpl, err := template.ParseFiles("templates/modules.html")
	if err != nil {
		err = ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		if err != nil {
			return
		}
		return
	}

	data := struct {
		Modules []module.Module
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
		params   module.NewModuleParams
	)

	if err = ctx.BodyParser(&params); err != nil {
		return err
	}

	if params.Title == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Title is required")
	}
	if params.Level == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Level is required")
	}
	if len(*params.Task) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Task is required")
	}

	err = h.UC.GenerateModule(context_, params)
	if err != nil {
		return err
	}

	return nil
}

func (h *ModuleHandler) GetNewModuleForm(ctx *fiber.Ctx) error {
	return ctx.Render("templates/create_module.html", fiber.Map{})
}
