package http

import "github.com/gofiber/fiber/v2"

func UserRoutes(app *fiber.App, h *UserHandler) {
	userGroup := app.Group("/user")
	userGroup.Get("/all", h.GetAll)

	adminGroup := userGroup.Group("/admin")
	adminGroup.Post("/sign_in", h.AdminSignIn)
	adminGroup.Post("/sign_up", h.AdminSignUp)
}
