package http

import "github.com/gofiber/fiber/v2"

func LearnRoutes(app *fiber.App, h *LearnHandler) {
	learnGroup := app.Group("/learn")

	ruleGroup := learnGroup.Group("/rule")
	ruleGroup.Get("/all", h.AllRules)
	ruleGroup.Post("/new", h.AddRule)
	ruleGroup.Get("/new", h.GetNewRuleForm)
}
