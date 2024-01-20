package user

import (
	"context"
	"english_bot_admin/internal/models"
)

type Repository interface {
	Select(ctx context.Context) ([]models.User, error)
}
