package http

import (
	"github.com/gofiber/fiber/v2"
)

func TaskRoutes(app *fiber.App, h *TaskHandler) {
	taskGroup := app.Group("/task")
	taskGroup.Get("/all", h.GetTasks)
	taskGroup.Post("/new", h.CreateTask)
	taskGroup.Post("/edit", h.EditTask)
}
