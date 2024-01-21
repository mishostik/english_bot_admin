package http

import (
	"english_bot_admin/internal/learning"
	"english_bot_admin/internal/models"
	"github.com/gofiber/fiber/v2"
)

type LearnHandler struct {
	UC learning.Usecase
}

func NewLearnHandler(UC learning.Usecase) *LearnHandler {
	return &LearnHandler{
		UC: UC,
	}
}

func (h *LearnHandler) AddRule(ctx *fiber.Ctx) error {
	var (
		params   models.NewRuleParams
		context_                       = ctx.Context()
		response *models.ResponseModel = &models.ResponseModel{}
		err      error
	)

	if err = ctx.BodyParser(&params); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	err = h.UC.AddRule(context_, &params)
	if err != nil {
		response.Error = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}
	return ctx.SendStatus(fiber.StatusOK)
}

func (h *LearnHandler) AllRules(ctx *fiber.Ctx) error {
	var (
		context_                            = ctx.Context()
		response *models.RulesResponseModel = &models.RulesResponseModel{}
	)
	rules, err := h.UC.GetAllRules(context_)
	if err != nil {
		response.Error = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Success = true
	response.Data = rules
	return ctx.Status(fiber.StatusInternalServerError).JSON(response)
}
