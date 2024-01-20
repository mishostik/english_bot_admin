package http

import (
	"bytes"
	"english_bot_admin/internal/models"
	"english_bot_admin/internal/user"
	"github.com/gofiber/fiber/v2"
	"html/template"
)

type UserHandler struct {
	uc user.Usecase
}

func NewUserHandler(ucase user.Usecase) *UserHandler {
	return &UserHandler{
		uc: ucase,
	}
}

func renderUsers(ctx *fiber.Ctx, users []models.User) {
	tmpl, err := template.ParseFiles("templates/users.html")
	if err != nil {
		err = ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		if err != nil {
			return
		}
		return
	}

	data := struct {
		Users  []models.User
		Amount int
	}{
		Users:  users,
		Amount: len(users),
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

func (h *UserHandler) GetAll(ctx *fiber.Ctx) error {
	context_ := ctx.Context()
	users, err := h.uc.GetAll(context_)
	if err != nil {
		return err
	}

	renderUsers(ctx, users)
	return nil
}

func (h *UserHandler) AdminSignIn(ctx *fiber.Ctx) error {
	var (
		params   models.AdminSignInParams
		context_ = ctx.Context()
		response = map[string]interface{}{"success": false, "error": ""}
		isExist  bool

		err error
	)

	if err = ctx.BodyParser(&params); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	isExist, err = h.uc.AdminSignIn(context_, &params)
	if err != nil {
		response["success"] = false
		response["error"] = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response["success"] = isExist
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (h *UserHandler) AdminSignUp(ctx *fiber.Ctx) error {
	var (
		params   models.AdminSignInParams
		context_ = ctx.Context()
		response = map[string]interface{}{"success": false, "error": ""}

		err error
	)

	if err = ctx.BodyParser(&params); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	err = h.uc.AdminSignUp(context_, &params)
	if err != nil {
		response["error"] = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	return ctx.SendStatus(fiber.StatusOK)
}
