package http

import (
	"github.com/gofiber/fiber/v2"
)

func TaskRoutes(app *fiber.App, h *TaskHandler) {
	taskGroup := app.Group("/task")
	taskGroup.Get("/all", h.GetTasks)
	taskGroup.Post("/new", h.CreateTask)
	taskGroup.Put("/edit/:id", h.EditTask)
	//taskGroup.Delete("/remove:/id", h.DeleteTask)
}
