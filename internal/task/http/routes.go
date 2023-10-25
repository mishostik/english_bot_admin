package http

import (
	"github.com/gofiber/fiber/v2"
)

func TaskRoutes(app *fiber.App, h *TaskHandler) {
	taskGroup := app.Group("/task")
	taskGroup.Get("/all", h.GetTasks)
	taskGroup.Get("/:id", h.GetTaskByID)
	taskGroup.Post("/new", h.CreateTask)
	taskGroup.Put("/update/:id", h.UpdateTask)
	taskGroup.Delete("/remove:/id", h.DeleteTask)
}
