package http

import (
	"english_bot_admin/internal/models"
	"english_bot_admin/internal/user"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	uc user.Usecase
}

func NewUserHandler(ucase user.Usecase) *UserHandler {
	return &UserHandler{
		uc: ucase,
	}
}

func (h *UserHandler) GetAll(ctx *fiber.Ctx) error {
	var (
		context_                            = ctx.Context()
		response *models.UsersResponseModel = &models.UsersResponseModel{}
	)
	users, err := h.uc.GetAll(context_)
	if err != nil {
		response.Error = err.Error()
		return ctx.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response.Success = true
	response.Data = users
	return ctx.Status(fiber.StatusOK).JSON(response)
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
