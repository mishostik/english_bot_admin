package http

import (
	"github.com/gofiber/fiber/v2"
)

func TaskRoutes(app *fiber.App, h *TaskHandler) {
	taskGroup := app.Group("/task")
	taskGroup.Get("/all", h.GetTasks)
	taskGroup.Post("/new", h.CreateTask)
	taskGroup.Post("/edit/:id", h.EditTask)
	//taskGroup.Delete("/remove:/id", h.DeleteTask)

	taskGroup.Get("/new", h.GetNewTaskForm)
	taskGroup.Get("/edit/:id", h.GetEditTaskForm)

	taskGroup.Post("/add_to_module", h.AddToModule)

	app.Get("", h.BaseView)

	app.Get("/access", h.GetTaskAccess)
	app.Post("/access", h.EditTaskAccess)

}
