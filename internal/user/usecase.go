package user

import (
	"context"
	"english_bot_admin/internal/models"
)

type Usecase interface {
	GetAll(context_ context.Context) ([]models.User, error)
	AdminSignUp(context_ context.Context, params *models.AdminSignInParams) error
	AdminSignIn(context_ context.Context, params *models.AdminSignInParams) (bool, error)
}
