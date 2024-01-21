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

// AddRule
// @Summary AddRule
// @Description Add rule
// @ID add-rule
// @Produce json
// @Param request body models.NewRuleParams true
// @Success 200 {object} models.ResponseModel
// @Failure 500 {object} models.ResponseModel
// @Router /learn/rule/new [post]
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

// AllRules
// @Summary AllRules
// @Description All rules
// @ID all-rules
// @Produce json
// @Success 200 {object} models.RulesResponseModel
// @Failure 500 {object} models.RulesResponseModel
// @Router /learn/rule/all [get]
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
