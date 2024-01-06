package http

import "github.com/gofiber/fiber/v2"

func LearnRoutes(app *fiber.App, h *LearnHandler) {
	moduleGroup := app.Group("/learn")
	moduleGroup.Get("/rule/all", h.AllRules)
	moduleGroup.Post("/rule/add", h.AddRule)
}
