package http

import "github.com/gofiber/fiber/v2"

func ModuleRoutes(app *fiber.App, h *ModuleHandler) {
	moduleGroup := app.Group("/module")
	moduleGroup.Get("/all", h.GetAllModules)
	moduleGroup.Post("/create", h.CreateModule)
}
