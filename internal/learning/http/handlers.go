package http

import (
	"bytes"
	"english_bot_admin/internal/learning"
	"github.com/gofiber/fiber/v2"
	"html/template"
)

type LearnHandler struct {
	UC learning.Usecase
}

func NewLearnHandler(UC learning.Usecase) *LearnHandler {
	return &LearnHandler{
		UC: UC,
	}
}

func (h *LearnHandler) renderRules(ctx *fiber.Ctx, rules []learning.Rule) {
	tmpl, err := template.ParseFiles("templates/rules.html")
	if err != nil {
		err = ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		if err != nil {
			return
		}
		return
	}

	data := struct {
		Rules []learning.Rule
	}{
		Rules: rules,
	}

	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, data); err != nil {
		err = ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		if err != nil {
			return
		}
		return
	}
	ctx.Set("Content-Type", "text/html")
	err = ctx.Status(fiber.StatusOK).Send(buf.Bytes())
	if err != nil {
		return
	}
}

func (h *LearnHandler) AddRule(ctx *fiber.Ctx) error {
	var (
		params   learning.NewRuleParams
		context_ = ctx.Context()
		err      error
	)

	if err = ctx.BodyParser(&params); err != nil {
		return err // todo ctx
	}

	err = h.UC.AddRule(context_, &params)
	if err != nil {
		return err // todo ctx
	}
	return nil
}

func (h *LearnHandler) AllRules(ctx *fiber.Ctx) error {
	context_ := ctx.Context()
	rules, err := h.UC.GetAllRules(context_)
	if err != nil {
		return err
	}
	h.renderRules(ctx, rules)
	return nil
}
