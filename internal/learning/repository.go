package learning

import (
	"context"
	"english_bot_admin/internal/models"
)

type Repository interface {
	InsertRule(ctx context.Context, rule *models.Rule) error
	SelectRules(ctx context.Context) ([]models.Rule, error)
}
