package user

import (
	"context"
	"english_bot_admin/internal/models"
)

type Repository interface {
	Select(ctx context.Context) ([]models.User, error)
	AdminVerification(ctx context.Context, params *models.AdminSignInParams) (bool, error)
	InsertAdmin(ctx context.Context, newAdmin *models.Admin) error
}
