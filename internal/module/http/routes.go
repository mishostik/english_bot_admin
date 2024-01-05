package http

import "github.com/gofiber/fiber/v2"

func ModuleRoutes(app *fiber.App, h *ModuleHandler) {
	moduleGroup := app.Group("/module")
	moduleGroup.Get("/all", h.GetAllModules)
	moduleGroup.Post("/new", h.CreateModule)
	moduleGroup.Get("/new", h.GetNewModuleForm)

	taskInnerGroup := moduleGroup.Group("/task")

	//taskInnerGroup.Get("/by_lvl", h.GetTasksByLvl)

	taskInnerGroup.Post("/by_lvl", h.GetTasksByLvl)

	// some questions ???
	taskInnerGroup.Post("/add", h.AddTasksByLvl)
}
